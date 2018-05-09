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

func (self *PlanJoinNode) GetNodeType() PlanNodeType {
	return JOINNODE
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
