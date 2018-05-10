package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanCombineNode struct {
	Inputs   []PlanNode
	Output   PlanNode
	Metadata *Util.Metadata
}

func NewPlanCombineNode(inputs []PlanNode, output PlanNode) *PlanCombineNode {
	return &PlanCombineNode{
		Inputs:   inputs,
		Output:   output,
		Metadata: Util.NewDefaultMetadata(),
	}
}

func (self *PlanCombineNode) GetNodeType() PlanNodeType {
	return COMBINENODE
}

func (self *PlanCombineNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanCombineNode) SetMetadata() (err error) {
	for _, input := range self.Inputs {
		if err = input.SetMetadata(); err != nil {
			return err
		}
		self.ColumnNames = append(self.ColumnNames, input.GetMetadata().ColumnNames...)
		self.ColumnTypes = append(self.ColumnTypes, input.GetMetadata().ColumnTypes...)
	}
	return nil
}

func (self *PlanCombineNode) String() string {
	res := "PlanCombineNode {\n"
	for _, n := range self.Inputs {
		res += n.String()
	}
	res += "}\n"
	return res
}
