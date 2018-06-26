package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Split"
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

	rbWriters := make([]*Split.SplitBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Split.NewSplitBuffer(enode.Metadata, nil, writer)
	}

	defer func() {
		for _, rbWriter := range rbWriters {
			rbWriter.Flush()
		}
	}()

	//send
	//no partitions
	jobs := make(chan *Split.Split)
	done := make(chan bool)
	k := 0

	for i := 0; i < int(Config.Conf.Runtime.ParallelNumber); i++ {
		go func() {
			defer func() {
				done <- true
			}()

			for {
				sp, ok := <-jobs
				if ok {
					for i := 0; i < sp.GetRowsNumber(); i++ {
						flag := true
						for _, filter := range enode.Filters {
							if ok, err := filter.Result(sp, i); !ok.(bool) || err != nil {
								flag = false
								break
							} else if err != nil {
								flag = false
								break
							}
						}

						if flag {
							if err := rbWriters[k%ln].Write(sp, i); err != nil {
								continue //should add err handler
							}
							k++
							k = k % ln
						}
					}

				} else {
					break
				}
			}
		}()
	}

	if !enode.PartitionInfo.IsPartition() {
		var sp *Split.Split
		for _, file := range enode.PartitionInfo.GetNoPartititonFiles() {
			reader := connector.GetReader(file, inputMetadata)
			//log.Println("[executor.scan]=====file", file)
			if err != nil {
				break
			}
			for err == nil {
				sp, err = reader(colIndexes)
				if err == io.EOF {
					err = nil
					break
				}
				if err != nil {
					break
				}
				jobs <- sp
			}
		}

	} else { //partitioned
		parColNum := enode.PartitionInfo.GetPartitionColumnNum()
		totColNum := inputMetadata.GetColumnNumber()
		dataColNum := totColNum - parColNum
		dataCols, parCols := []int{}, []int{}
		var sp *Split.Split
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
			par := enode.PartitionInfo.GetPartition(i)
			for _, file := range enode.PartitionInfo.GetPartitionFiles(i) {
				reader := connector.GetReader(file, inputMetadata)
				//log.Println("======", self.Name, file)
				if err != nil {
					break
				}
				for err == nil {
					sp, err = reader(dataCols)
					if err == io.EOF {
						err = nil
						break
					}
					if err != nil {
						break
					}

					sp.Metadata = enode.Metadata
					rn := sp.GetRowsNumber()
					for _, index := range parCols {
						parCol := make([]interface{}, rn)
						for i := 0; i < rn; i++ {
							parCol[i] = par[index]
						}
						sp.AppendColumn(parCol)
					}

					jobs <- sp
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
