package Plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanSelectNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *Util.Metadata
	SelectItems []*SelectItemNode
	IsAggregate bool
}

func NewPlanSelectNode(input PlanNode, items []parser.ISelectItemContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		Metadata:    Util.NewMetadata(),
		SelectItems: []*SelectItemNode{},
	}
	for i := 0; i < len(items); i++ {
		itemNode := NewSelectItemNode(items[i])
		res.SelectItems = append(res.SelectItems, itemNode)
		if itemNode.IsAggregate() {
			res.IsAggregate = true
		}
	}
	return res
}

func (self *PlanSelectNode) GetNodeType() PlanNodeType {
	return SELECTNODE
}

func (self *PlanSelectNode) GetInputs() []PlanNode {
	return []PlanNode{self.Input}
}

func (self *PlanSelectNode) SetInputs(inputs []PlanNode) {
	self.Input = inputs[0]
}

func (self *PlanSelectNode) GetOutput() PlanNode {
	return self.Output
}

func (self *PlanSelectNode) SetOutput(output PlanNode) {
	self.Output = output
}

func (self *PlanSelectNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanSelectNode) SetMetadata() error {
	self.Input.SetMetadata()
	md := self.Input.GetMetadata()
	colNames, colTypes := []string{}, []Util.Type{}
	for _, item := range self.SelectItems {
		names, types, err := item.GetNamesAndTypes(md)
		if err != nil {
			return err
		}
		colNames = append(colNames, names...)
		colTypes = append(colTypes, types...)
	}

	if len(colNames) != len(colTypes) {
		return fmt.Errorf("length error")
	}
	self.Metadata = Util.NewMetadata()
	for i, name := range colNames {
		t := colTypes[i]
		column := Util.NewColumnMetadata(t, strings.Split(name, ".")...)
		self.Metadata.AppendColumn(column)
	}
	self.Metadata.Reset()

	return nil
}

func (self *PlanSelectNode) String() string {
	res := "PlanSelectNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "Metadata: " + fmt.Sprint(self.Metadata) + "\n"
	res += "SelectItems: " + fmt.Sprint(self.SelectItems) + "\n"
	res += "}\n"
	return res
}
