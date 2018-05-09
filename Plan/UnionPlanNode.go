package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type UnionType int32

const (
	_ UnionType = iota
	INTERSECT
	UNION
	EXCEPT
)

type PlanUnionNode struct {
	LeftInput  PlanNode
	RightInput PlanNode
	Operator   UnionType
	Metadata   *Util.Metadata
}

func NewPlanUnionNode(left PlanNode, right PlanNode, op antlr.Token) *PlanUnionNode {
	var operator UnionType
	switch op.GetText() {
	case "INTERSECT":
		operator = INTERSECT
	case "UNION":
		operator = UNION
	case "EXCEPT":
		operator = EXCEPT
	}

	res := &PlanUnionNode{
		LeftInput:  left,
		RightInput: right,
		Operator:   operator,
		Metadata:   Util.NewDefaultMetadata(),
	}
	return res
}

func (self *PlanUnionNode) GetNodeType() PlanNodeType {
	return UNIONNODE
}

func (self *PlanUnionNode) String() string {
	res := "PlanUnionNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "Operator: " + fmt.Sprint(self.Operator) + "\n"
	res += "}\n"
	return res
}
