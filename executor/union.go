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

func (self *Executor) SetInstructionUnion(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanUnionNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.LeftInput, &enode.RightInput}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunUnion() (err error) {
	fname := fmt.Sprintf("executor_%v_scan_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		self.AddLogInfo(err, pb.LogLevel_ERR)
		self.Clear()
	}()
	writer := self.Writers[0]
	//enode := self.EPlanNode.(*EPlan.EPlanUnionNode)

	//read md
	if len(self.Readers) != 2 {
		return fmt.Errorf("union readers number %v <> 2", len(self.Readers))
	}

	md := &metadata.Metadata{}
	if len(self.Readers) != 2 {
		return fmt.Errorf("union input number error")
	}
	for _, reader := range self.Readers {
		if err = util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	//write md
	if err = util.WriteObject(writer, md); err != nil {
		return err
	}

	rbWriter := row.NewRowsBuffer(md, nil, writer)
	defer func() {
		rbWriter.Flush()
	}()

	//write rows
	var rg *row.RowsGroup
	for _, reader := range self.Readers {
		rbReader := row.NewRowsBuffer(md, reader, nil)
		for {
			rg, err = rbReader.Read()
			if err == io.EOF {
				err = nil
				break
			}
			if err != nil {
				return err
			}
			if err = rbWriter.Write(rg); err != nil {
				return err
			}
		}
	}

	logger.Infof("RunUnion finished")
	return err
}
