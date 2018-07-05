package Executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionGroupByLocal(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByLocalNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunGroupByLocal() (err error) {
	Logger.Infof("RunGroupByLocal")
	fname := fmt.Sprintf("executor_%v_groupbylocal_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer self.Clear()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByLocalNode)

	reader := self.Readers[0]
	writer := self.Writers[0]
	md := &Metadata.Metadata{}

	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReader := Row.NewRowsBuffer(md, reader, nil)
	rbWriter := Row.NewRowsBuffer(md, nil, writer)

	//group by
	var row *Row.Row
	var rgs = make(map[string]*Row.RowsGroup)
	for {
		row, err = rbReader.ReadRow()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		key := row.GetKeyString()
		if _, ok := rgs[key]; !ok {
			rgs[key] = Row.NewRowsGroup(enode.Metadata)
		}
		rgs[key].Write(row)
	}

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	for _, rg := range rgs {
		var (
			ok  interface{} = true
			err error       = nil
		)
		if enode.GroupBy.Having != nil {
			ok, err = enode.GroupBy.Having.Result(rg)
		}

		if err == nil && ok.(bool) {
			rg.Reset()
			for {
				row, err := rg.Read()
				if err == io.EOF {
					err = nil
					break
				}
				if err != nil {
					return err
				}
				if err = rbWriter.WriteRow(row); err != nil {
					return err
				}
			}
		} else if err != nil {
			return err
		}
	}

	return err
}
