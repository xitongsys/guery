package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionGroupBy(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanGroupByNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunGroupBy() (err error) {
	Logger.Infof("RunGroupBy")
	fname := fmt.Sprintf("executor_%v_groupby_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	if self.Instruction == nil {
		return fmt.Errorf("no instruction")
	}
	enode := self.EPlanNode.(*EPlan.EPlanGroupByNode)

	md := &Metadata.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = Util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	enode.Metadata.ClearKeys()
	enode.Metadata.AppendKeyByType(Type.STRING)
	if err = Util.WriteObject(writer, enode.Metadata); err != nil {
		return err
	}

	rbReader := Row.NewRowsBuffer(md, reader, nil)
	rbWriter := Row.NewRowsBuffer(enode.Metadata, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//group by

	if err := enode.GroupBy.Init(md); err != nil {
		return err
	}
	var rg *Row.RowsGroup
	for {
		rg, err = rbReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		keys, err := enode.GroupBy.Result(rg)
		if err != nil {
			return err
		}
		rg.AppendKeyColumns(keys)

		if err := rbWriter.Write(rg); err != nil {
			return err
		}
	}

	return err
}
