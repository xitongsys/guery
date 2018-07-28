package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionLimit(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanLimitNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunLimit() (err error) {
	fname := fmt.Sprintf("executor_%v_limit_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	enode := self.EPlanNode.(*eplan.EPlanLimitNode)
	writer := self.Writers[0]
	md := &metadata.Metadata{}
	//read md
	for _, reader := range self.Readers {
		if err = util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	if err = util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReaders := make([]*row.RowsBuffer, len(self.Readers))
	for i, reader := range self.Readers {
		rbReaders[i] = row.NewRowsBuffer(md, reader, nil)
	}
	rbWriter := row.NewRowsBuffer(md, nil, writer)

	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var rg *row.RowsGroup
	readRowCnt := int64(0)
	for _, rbReader := range rbReaders {
		for readRowCnt < *(enode.LimitNumber) {
			rg, err = rbReader.Read()
			if err == io.EOF || readRowCnt >= *(enode.LimitNumber) {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			readRowCnt += int64(rg.GetRowsNumber())
			if err = rbWriter.Write(rg); err != nil {
				return err
			}
		}
	}

	logger.Infof("RunAggregate finished")
	return nil
}
