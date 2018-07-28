package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/connector"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/plan"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionShow(instruction *pb.Instruction) error {
	logger.Infof("set instruction show")
	var enode eplan.EPlanShowNode
	var err error
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}

	self.EPlanNode = &enode
	self.Instruction = instruction
	self.InputLocations = []*pb.Location{}
	self.OutputLocations = append(self.OutputLocations, &enode.Output)
	return nil
}

func (self *Executor) RunShow() (err error) {
	fname := fmt.Sprintf("executor_%v_show_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
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

	enode := self.EPlanNode.(*eplan.EPlanShowNode)
	ctr, err := connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	md := enode.Metadata
	writer := self.Writers[0]
	//write metadata
	if err = util.WriteObject(writer, md); err != nil {
		return err
	}

	rbWriter := row.NewRowsBuffer(md, nil, writer)

	var showReader func() (*row.Row, error)
	//writer rows
	switch enode.ShowType {
	case plan.SHOWCATALOGS:
	case plan.SHOWSCHEMAS:
		showReader = ctr.ShowSchemas(enode.Catalog, enode.Schema, enode.Table, enode.LikePattern, enode.Escape)
	case plan.SHOWTABLES:
		showReader = ctr.ShowTables(enode.Catalog, enode.Schema, enode.Table, enode.LikePattern, enode.Escape)
	case plan.SHOWCOLUMNS:
		showReader = ctr.ShowColumns(enode.Catalog, enode.Schema, enode.Table)
	case plan.SHOWPARTITIONS:
		showReader = ctr.ShowPartitions(enode.Catalog, enode.Schema, enode.Table)
	}

	for {
		r, err := showReader()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}

		if err = rbWriter.WriteRow(r); err != nil {
			return err
		}
	}

	if err = rbWriter.Flush(); err != nil {
		return err
	}

	logger.Infof("RunShowTables finished")
	return err

}
