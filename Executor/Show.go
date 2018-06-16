package Executor

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/Connector"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Plan"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionShow(instruction *pb.Instruction) error {
	Logger.Infof("set instruction show")
	var enode EPlan.EPlanShowNode
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

	enode := self.EPlanNode.(*EPlan.EPlanShowNode)
	connector, err := Connector.NewConnector(enode.Catalog, enode.Schema, enode.Table)
	if err != nil {
		return err
	}

	md := enode.Metadata
	writer := self.Writers[0]

	//write metadata
	if err = Util.WriteObject(writer, md); err != nil {
		return err
	}

	rbWriter := Row.NewRowsBuffer(md, nil, writer)

	var showReader func() (*Row.Row, error)
	//writer rows
	switch enode.ShowType {
	case Plan.SHOWCATALOGS:
	case Plan.SHOWSCHEMAS:
		showReader = connector.ShowSchemas(enode.LikePattern, enode.Escape)
	case Plan.SHOWTABLES:
		showReader = connector.ShowTables(enode.Schema, enode.LikePattern, enode.Escape)
	}

	for {
		row, err := showReader()
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

	if err = rbWriter.Flush(); err != nil {
		return err
	}

	Logger.Infof("RunShowTables finished")
	return err

}
