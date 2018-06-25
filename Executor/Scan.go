package Executor

import (
	"fmt"
	"io"
	"log"

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

	//send rows
	//no partitions
	jobs := make(chan *Row.Row)
	done := make(chan bool)
	k := 0

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		go func() {
			defer func() {
				done <- true
			}()

			for {
				row, ok := <-jobs

				if ok {
					rg := Row.NewRowsGroup(enode.Metadata)
					rg.Write(row)
					flag := true
					for _, filter := range enode.Filters {
						rg.Reset()
						if ok, err := filter.Result(rg); !ok.(bool) || err != nil {
							flag = false
							break
						} else if err != nil {
							flag = false
							break
						}
					}

					if flag {
						if err := rbWriters[k%ln].WriteRow(row); err != nil {
							continue //should add err handler
						}
						k++
						k = k % ln
					}

				} else {
					break
				}
			}
		}()
	}

	if !enode.PartitionInfo.IsPartition() {
		var row *Row.Row
		for _, file := range enode.PartitionInfo.GetNoPartititonFiles() {
			reader := connector.GetReader(file, inputMetadata)
			//log.Println("[executor.scan]=====file", file)
			if err != nil {
				break
			}
			for err == nil {
				row, err = reader(colIndexes)
				if err == io.EOF {
					err = nil
					break
				}
				if err != nil {
					break
				}

				jobs <- row

			}
		}

	} else { //partitioned
		parColNum := enode.PartitionInfo.GetPartitionColumnNum()
		totColNum := inputMetadata.GetColumnNumber()
		dataColNum := totColNum - parColNum
		dataCols, parCols := []int{}, []int{}
		var row *Row.Row
		for _, index := range colIndexes {
			if index < dataColNum {
				dataCols = append(dataCols, index) //column from input
			} else {
				parCols = append(parCols, index-dataColNum) //column from partition
			}
		}
		for i := totColNum - 1; i >= dataColNum; i-- {
			inputMetadata.DeleteColumnByIndex(i)
		}

		for i := 0; i < enode.PartitionInfo.GetPartitionNum(); i++ {
			for _, file := range enode.PartitionInfo.GetPartitionFiles(i) {
				reader := connector.GetReader(file, inputMetadata)
				log.Println("======", self.Name, file, err, inputMetadata)
				if err != nil {
					break
				}
				for err == nil {
					row, err = reader(dataCols)
					//log.Println("======", err, dataCols, row)
					if err == io.EOF {
						err = nil
						break
					}
					if err != nil {
						break
					}

					parRow := enode.PartitionInfo.GetPartitionRow(i)
					for _, index := range parCols {
						//log.Println("=====", parRow.Vals[index], reflect.TypeOf(parRow.Vals[index]))
						row.AppendVals(parRow.Vals[index])
					}

					jobs <- row
				}
			}
		}
	}
	close(jobs)
	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		<-done
	}

	Logger.Infof("RunScan finished")
	return err
}
