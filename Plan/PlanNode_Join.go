package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type JoinType int32

const (
	_ JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	INNERJOIN
)

type PlanJoinNode struct {
	Metadata              *Util.Metadata
	LeftInput, RightInput PlanNode
	Output                PlanNode
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
}

func NewPlanJoinNode(leftInput PlanNode, rightInput PlanNode, joinType JoinType, joinCriteria *JoinCriteriaNode) *PlanJoinNode {
	res := &PlanJoinNode{
		Metadata:     Util.NewDefaultMetadata(),
		LeftInput:    leftInput,
		RightInput:   rightInput,
		JoinType:     joinType,
		JoinCriteria: joinCriteria,
	}
	return res
}

func (self *PlanJoinNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanJoinNode) GetNodeType() PlanNodeType {
	return JOINNODE
}

func (self *PlanJoinNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanJoinNode) SetMetadata() (err error) {
	if err = self.LeftInput.SetMetadata(); err != nil {
		return err
	}
	if err = self.RightInput.SetMetadata(); err != nil {
		return err
	}

	mdl, mdr := self.LeftInput.GetMetadata(), self.RightInput.GetMetadata()
	self.Metadata.ColumnNames = append(self.Metadata.ColumnNames, mdl.ColumnNames...)
	self.Metadata.ColumnNames = append(self.Metadata.ColumnNames, mdr.ColumnNames...)
	self.Metadata.ColumnTypes = append(self.Metadata.ColumnTypes, mdl.ColumnTypes...)
	self.Metadata.ColumnTypes = append(self.Metadata.ColumnTypes, mdr.ColumnTypes...)
	self.Metadata.Reset()
	return nil
}

func (self *PlanJoinNode) String() string {
	res := "PlanJoinNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "JoinType: " + fmt.Sprint(self.JoinType) + "\n"
	res += "JoinCriteria: " + fmt.Sprint(self.JoinCriteria) + "\n"
	res += "}\n"
	return res
}
