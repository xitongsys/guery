package Executor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionGroupBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByNode
	if err = gob.NewDecoder(bytes.NewBuffer(instruction.EncodedEPlanNodeBytes)).Decode(&enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for _, loc := range enode.Inputs {
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{}
	for _, loc := range enode.Outputs {
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func (self *Executor) RunGroupBy() (err error) {
	Logger.Infof("RunGroupBy")
	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByNode)

	mds := make([]*Util.Metadata, len(self.Readers))
	for i, reader := range self.Readers {
		mds[i] = &Util.Metadata{}
		if err = Util.ReadObject(reader, mds[i]); err != nil {
			return err
		}
	}

	//write metadata
	gmd := enode.Metadata
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, gmd); err != nil {
			return err
		}
	}

	//group by
	var row *Util.Row
	var rowsBufs = make(map[string]*Util.RowsBuffer)
	for i, reader := range self.Readers {
		for {
			row, err = Util.ReadRow(reader)
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				break
			}

			row.Key, err = self.CalGroupByKey(enode, mds[i], row)
			if err != nil {
				return err
			}
			if _, ok := rowsBufs[row.Key]; !ok {
				rowsBufs[row.Key] = Util.NewRowsBuffer(gmd)
			}
			rowsBufs[row.Key].Write(row)
		}
	}

	//write rows
	Done := make(chan int)
	ErrChan, TaskChan := make(chan error, len(rowsBufs)), make(chan *Util.RowsBuffer)
	for i := 0; i < len(self.Writers); i++ {
		go func(wi int) {
			writer := self.Writers[wi]
			for {
				select {
				case <-Done:
					return
				case rb := <-TaskChan:
					for {
						row, err := rb.Read()
						if err != nil {
							ErrChan <- err
							break
						}
						Util.WriteRow(writer, row)
					}
				}
			}
		}(i)
	}

	for _, rb := range rowsBufs {
		TaskChan <- rb
	}

	for i := 0; i < len(rowsBufs); i++ {
		e := <-ErrChan
		if e != nil && e != io.EOF {
			err = e
		}
	}
	close(Done)
	return err
}

func (self *Executor) CalGroupByKey(enode *EPlan.EPlanGroupByNode, md *Util.Metadata, row *Util.Row) (string, error) {
	rowsBuf := Util.NewRowsBuffer(md)
	rowsBuf.Write(row)
	res, err := enode.GroupBy.Result(rowsBuf)
	if err != nil {
		return res, err
	}
	return res, nil
}
