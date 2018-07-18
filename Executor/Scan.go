package Executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionScan(instruction *pb.Instruction) error {
	Logger.Infof("set instruction scan")

	var enode EPlan.EPlanScanNode
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
			Util.WriteEOFMessage(self.Writers[i])
			self.Writers[i].(io.WriteCloser).Close()
		}
		self.Clear()

	}()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}

	enode := self.EPlanNode.(*EPlan.EPlanScanNode)
	connector, err := Connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	ln := len(self.Writers)
	//send metadata
	for i := 0; i < ln; i++ {
		if err = Util.WriteObject(self.Writers[i], enode.Metadata); err != nil {
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

	rbWriters := make([]*Row.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Row.NewRowsBuffer(enode.Metadata, nil, writer)
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
	jobs := make(chan *Row.RowsGroup)
	var wg sync.WaitGroup

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
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
							break //should add err handler
						}
						flags := flagsi.([]interface{})
						rgtmp := Row.NewRowsGroup(enode.Metadata)
						for i, f := range flags {
							if f.(bool) {
								rgtmp.AppendVals(rg.GetRowVals(i)...)
							}
						}
						rg = rgtmp
					}

					if err := rbWriters[k].Write(rg); err != nil {
						continue //should add err handler
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
		var rg *Row.RowsGroup
		for _, file := range enode.PartitionInfo.GetNoPartititonFiles() {
			reader := connector.GetReader(file, inputMetadata)
			//log.Println("[executor.scan]=====file", file)
			if err != nil {
				break
			}
			for err == nil {
				rg, err = reader(colIndexes)
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
				parCols = append(parCols, index) //column from partition
			}
		}
		parMD := inputMetadata.SelectColumnsByIndexes(parCols)

		for i := totColNum - 1; i >= dataColNum; i-- {
			inputMetadata.DeleteColumnByIndex(i)
		}

		for i := 0; i < enode.PartitionInfo.GetPartitionNum(); i++ {
			for _, file := range enode.PartitionInfo.GetPartitionFiles(i) {
				reader := connector.GetReader(file, inputMetadata)
				//log.Println("======", self.Name, file)
				if err != nil {
					break
				}
				for err == nil {
					dataRG, err := reader(dataCols)
					//log.Println("======", err, dataCols, row)
					if err == io.EOF {
						err = nil
						break
					}
					if err != nil {
						break
					}

					parRow := enode.PartitionInfo.GetPartitionRow(i)
					parRG := Row.NewRowsGroup(parMD)
					for i := 0; i < dataRG.GetRowsNumber(); i++ {
						parRG.Write(parRow)
					}

					rg := Row.NewRowsGroup(enode.Metadata)
					rg.AppendColumns(dataRG.Vals...)
					rg.AppendColumns(parRG.Vals...)

					jobs <- rg
				}
			}
		}
	}
	close(jobs)
	wg.Wait()

	Logger.Infof("RunScan finished")
	return err
}
