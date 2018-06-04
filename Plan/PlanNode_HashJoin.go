package Plan

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
)

type PlanHashJoinNode struct {
	Metadata              *Util.Metadata
	LeftInput, RightInput PlanNode
	Output                PlanNode
	JoinType              JoinType
	JoinCriteria          *JoinCriteriaNode
	LeftKeys, RightKeys   []*ValueExpressionNode
}

func NewPlanHashJoinNodeFromJoinNode(node *PlanJoinNode, leftKeys, rightKeys []*ValueExpressionNode) *PlanHashJoinNode {
	return &PlanHashJoinNode{
		Metadata:     node.Metadata,
		LeftInput:    node.LeftInput,
		RightInput:   node.RightInput,
		JoinType:     node.JoinType,
		JoinCriteria: node.JoinCriteria,
		LeftKeys:     leftKeys,
		RightKeys:    rightKeys,
	}
}

func NewPlanHashJoinNode(leftInput PlanNode, rightInput PlanNode, joinType JoinType, joinCriteria *JoinCriteriaNode, leftKeys, rightKeys []*ValueExpressionNode) *PlanHashJoinNode {
	res := &PlanHashJoinNode{
		Metadata:     Util.NewMetadata(),
		LeftInput:    leftInput,
		RightInput:   rightInput,
		JoinType:     joinType,
		JoinCriteria: joinCriteria,
		LeftKeys:     leftKeys,
		RightKeys:    rightKeys,
	}
	return res
}

func (self *PlanHashJoinNode) GetInputs() []PlanNode {
	return []PlanNode{self.LeftInput, self.RightInput}
}

func (self *PlanHashJoinNode) SetInputs(inputs []PlanNode) {
	self.LeftInput, self.RightInput = inputs[0], inputs[1]
}

func (self *PlanHashJoinNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanHashJoinNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanHashJoinNode) GetNodeType() PlanNodeType {
	return JOINNODE
}

func (self *PlanHashJoinNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanHashJoinNode) SetMetadata() (err error) {
	if err = self.LeftInput.SetMetadata(); err != nil {
		return err
	}
	if err = self.RightInput.SetMetadata(); err != nil {
		return err
	}

	mdl, mdr := self.LeftInput.GetMetadata(), self.RightInput.GetMetadata()
	self.Metadata = Util.JoinMetadata(mdl, mdr)
	return nil
}

func (self *PlanHashJoinNode) String() string {
	res := "PlanHashJoinNode {\n"
	res += "LeftInput: " + self.LeftInput.String() + "\n"
	res += "RightInput: " + self.RightInput.String() + "\n"
	res += "JoinType: " + fmt.Sprint(self.JoinType) + "\n"
	res += "JoinCriteria: " + fmt.Sprint(self.JoinCriteria) + "\n"
	res += "}\n"
	return res
}
