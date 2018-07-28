package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/connector"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionScan(instruction *pb.Instruction) error {
	logger.Infof("set instruction scan")

	var enode eplan.EPlanScanNode
	var err error
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	enode.PartitionInfo.Decode() //partitioninfo must decode firstly

	self.EPlanNode = &enode
	self.Instruction = instruction
	for i := 0; i < len(enode.Outputs); i++ {
		loc := enode.Outputs[i]
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunScan() (err error) {
	fname := fmt.Sprintf("executor_%v_scan_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		for i := 0; i < len(self.Writers); i++ {
			util.WriteEOFMessage(self.Writers[i])
			self.Writers[i].(io.WriteCloser).Close()
		}
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()

	}()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*eplan.EPlanScanNode)
	ctr, err := connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	ln := len(self.Writers)
	//send metadata
	for i := 0; i < ln; i++ {
		if err = util.WriteObject(self.Writers[i], enode.Metadata); err != nil {
			return err
		}
	}

	colIndexes := []int{}
	inputMetadata := enode.InputMetadata
	for _, c := range enode.Metadata.Columns {
		cn := c.ColumnName
		index, err := inputMetadata.GetIndexByName(cn)
		if err != nil {
			return err
		}
		colIndexes = append(colIndexes, index)
	}

	rbWriters := make([]*row.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = row.NewRowsBuffer(enode.Metadata, nil, writer)
	}

	defer func() {
		for _, rbWriter := range rbWriters {
			rbWriter.Flush()
		}
	}()

	//init
	for _, f := range enode.Filters {
		if err := f.Init(enode.Metadata); err != nil {
			return err
		}
	}

	//send rows
	//no partitions
	jobs := make(chan *row.RowsGroup)
	var wg sync.WaitGroup

	for i := 0; i < int(config.Conf.Runtime.ParallelNumber); i++ {
		wg.Add(1)
		go func(ki int) {
			defer func() {
				wg.Done()
			}()

			k := ki % ln

			for {
				rg, ok := <-jobs
				if ok {
					for _, filter := range enode.Filters { //TODO: improve performance, add flag in RowsGroup?
						flagsi, err := filter.Result(rg)
						if err != nil {
							self.AddLogInfo(err, pb.LogLevel_ERR)
							break
						}
						flags := flagsi.([]interface{})
						rgtmp := row.NewRowsGroup(enode.Metadata)

						for i, f := range flags {
							if f.(bool) {
								rgtmp.AppendValRow(rg.GetRowVals(i)...)
							}
						}
						rg = rgtmp
					}

					if err := rbWriters[k].Write(rg); err != nil {
						self.AddLogInfo(err, pb.LogLevel_ERR)
						break
					}
					k++
					k = k % ln

				} else {
					break
				}
			}
		}(i)
	}

	if !enode.PartitionInfo.IsPartition() {
		for _, file := range enode.PartitionInfo.GetNoPartititonFiles() {
			reader := ctr.GetReader(file, inputMetadata)
			if err != nil {
				break
			}
			for err == nil {
				rg, err := reader(colIndexes)
				if err == io.EOF {
					err = nil
					break
				}
				if err != nil {
					break
				}

				jobs <- rg

			}
		}

	} else { //partitioned
		parColNum := enode.PartitionInfo.GetPartitionColumnNum()
		totColNum := inputMetadata.GetColumnNumber()
		dataColNum := totColNum - parColNum
		dataCols, parCols := []int{}, []int{}

		for _, index := range colIndexes {
			if index < dataColNum {
				dataCols = append(dataCols, index) //column from input
			} else {
				parCols = append(parCols, index-dataColNum) //column from partition
			}
		}
		parMD := inputMetadata.SelectColumnsByIndexes(parCols)

		for i := totColNum - 1; i >= dataColNum; i-- {
			inputMetadata.DeleteColumnByIndex(i)
		}

		for i := 0; i < enode.PartitionInfo.GetPartitionNum(); i++ {
			parFullRow := enode.PartitionInfo.GetPartitionRow(i)
			parRow := row.NewRow()
			for _, index := range parCols {
				parRow.AppendVals(parFullRow.Vals[index])
			}

			for _, file := range enode.PartitionInfo.GetPartitionFiles(i) {
				reader := ctr.GetReader(file, inputMetadata)
				if err != nil {
					break
				}
				for err == nil {
					dataRG, err := reader(dataCols)
					if err == io.EOF {
						err = nil
						break
					}
					if err != nil {
						break
					}

					parRG := row.NewRowsGroup(parMD)
					for i := 0; i < dataRG.GetRowsNumber(); i++ {
						parRG.Write(parRow)
					}

					rg := row.NewRowsGroup(enode.Metadata)
					rg.ClearColumns()
					rg.AppendValColumns(dataRG.Vals...)
					rg.AppendValColumns(parRG.Vals...)

					jobs <- rg
				}
			}
		}
	}
	close(jobs)
	wg.Wait()

	logger.Infof("RunScan finished")
	return err
}
