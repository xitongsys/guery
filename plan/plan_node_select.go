package Plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Config"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/parser"
)

type PlanSelectNode struct {
	Input       PlanNode
	Output      PlanNode
	Metadata    *Metadata.Metadata
	SelectItems []*SelectItemNode
	Having      *BooleanExpressionNode
	IsAggregate bool
}

func NewPlanSelectNode(runtime *Config.ConfigRuntime, input PlanNode, items []parser.ISelectItemContext, having parser.IBooleanExpressionContext) *PlanSelectNode {
	res := &PlanSelectNode{
		Input:       input,
		Metadata:    Metadata.NewMetadata(),
		SelectItems: []*SelectItemNode{},
		Having:      nil,
	}
	for i := 0; i < len(items); i++ {
		itemNode := NewSelectItemNode(runtime, items[i])
		res.SelectItems = append(res.SelectItems, itemNode)
		if itemNode.IsAggregate() {
			res.IsAggregate = true
		}
	}

	if having != nil {
		res.Having = NewBooleanExpressionNode(runtime, having)
		if res.Having.IsAggregate() {
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

func (self *PlanSelectNode) GetMetadata() *Metadata.Metadata {
	return self.Metadata
}

func (self *PlanSelectNode) SetMetadata() error {
	if err := self.Input.SetMetadata(); err != nil {
		return err
	}
	md := self.Input.GetMetadata()
	colNames, colTypes := []string{}, []Type.Type{}
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
	self.Metadata = Metadata.NewMetadata()
	for i, name := range colNames {
		t := colTypes[i]
		column := Metadata.NewColumnMetadata(t, strings.Split(name, ".")...)
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
