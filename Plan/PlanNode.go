package Plan

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type PlanNodeType int32

const (
	_ PlanNodeType = iota
	SCANNODE
	JOINNODE
	FILTERNODE
	ORDERBYNODE
	LIMITNODE
	SELECTNODE
	UNIONNODE
	HAVINGNODE
	UNITNODE
)

type PlanNode interface {
	Execute() *DataSource.DataSource
}

//////////////////
type PlanOrderByNode struct {
	Name  string
	Input PlanNode
}

func NewPlanOrderByNode(ctx *Context.Context, input PlanNode, sortItems []parser.ISortItemContext) *PlanOrderByNode {
	res := &PlanOrderByNode{
		Input: input,
	}
	return res
}

func (self *PlanOrderByNode) Execute() *DataSource.DataSource {
	return self.Input.Execute()
}

/////////////////
type PlanLimitNode struct {
	Name        string
	Input       PlanNode
	LimitNumber *int64
}

func NewPlanLimitNode(ctx *Context.Context, input PlanNode, t antlr.TerminalNode) *PlanLimitNode {
	res := &PlanLimitNode{
		Input: input,
	}
	if ns := t.GetText(); ns != "ALL" {
		var num int64
		fmt.Sscanf(ns, "%d", &num)
		res.LimitNumber = &num
	}
	return res
}

func (self *PlanLimitNode) Execute() *DataSource.DataSource {
	return self.Input.Execute()
}

////////////////
type PlanUnionNode struct {
	Name       string
	LeftInput  PlanNode
	RightInput PlanNode
	Operator   Common.Operator
}

func NewPlanUnionNode(ctx *Context.Context, name string, left PlanNode, right PlanNode, op antlr.Token) *PlanUnionNode {
	var operator Common.Operator
	switch op.GetText() {
	case "INTERSECT":
		operator = Common.INTERSECT
	case "UNION":
		operator = Common.UNION
	case "EXCEPT":
		operator = Common.EXCEPT
	}

	res := &PlanUnionNode{
		Name:       name,
		LeftInput:  left,
		RightInput: right,
		Operator:   operator,
	}
	return res
}

func (self *PlanUnionNode) Execute() *DataSource.DataSource {
	return self.LeftInput.Execute()
}

//////////////
type PlanFiliterNode struct {
	Name  string
	Input PlanNode
}

func NewPlanFiliterNode(ctx *Context.Context, input PlanNode, t parser.IBooleanExpressionContext) *PlanFiliterNode {
	res := &PlanFiliterNode{
		Input: input,
	}
	return res
}

func (self *PlanFiliterNode) Execute() *DataSource.DataSource {
	return self.Input.Execute()
}

//////////////
type PlanHavingNode struct {
	Name  string
	Input PlanNode
}

func NewPlanHavingNode(ctx *Context.Context, input PlanNode, t parser.IBooleanExpressionContext) *PlanHavingNode {
	res := &PlanHavingNode{
		Input: input,
	}
	return res
}

func (self *PlanHavingNode) Execute() *DataSource.DataSource {
	return self.Input.Execute()
}

////////////////
type PlanSelectNode struct {
	Name        string
	Input       PlanNode
	SelectItems []*SelectItemNode
	GroupBy     *GroupByNode
}

func NewPlanSelectNode(ctx *Context.Context, name string, input PlanNode, items []parser.ISelectItemContext, groupBy parser.IGroupByContext) *PlanSelectNode {

	res := &PlanSelectNode{
		Name:        name,
		Input:       input,
		SelectItems: []*SelectItemNode{},
		GroupBy:     NewGroupByNode(ctx, groupBy),
	}
	for i := 0; i < len(items); i++ {
		res.SelectItems = append(res.SelectItems, NewSelectItemNode(ctx, items[i]))
	}
	return res
}

func (self *PlanSelectNode) Execute() *DataSource.DataSource {
	var ds *DataSource.DataSource
	if self.Input != nil {
		ds = self.Input.Execute()
	}

	dss := []*DataSource.DataSource{}
	if self.GroupBy != nil {
		dsMap := make(map[string]*DataSource.DataSource)
		for i := 0; i < ds.GetRowNum(); i++ {
			dsr := ds.SelectRow()

			key := self.GroupBy.Result(dsr)

			if _, ok := dsMap[key]; !ok {
				dsMap[key] = dsr
			} else {
				dsMap[key].Merge(dsr)
			}
			ds.Next()
		}
		for _, val := range dsMap {
			dss = append(dss, val)
		}

	} else {
		dss = append(dss, ds)
	}

	size := len(dss)

	columnBuffers := make([]DataSource.ColumnBuffer, len(self.SelectItems))

	for i := 0; i < len(self.SelectItems); i++ {
		item := self.SelectItems[i]
		ci := len(columnBuffers)
		for j := 0; j < size; j++ {
			dss[j].Reset()
			res := item.Result(dss[j])
			cii := ci
			for k := 0; k < len(res); k++ {
				if ci >= len(columnBuffers) {
					columnBuffers = append(columnBuffers, DataSource.NewMemColumnBuffer())
				}
				columnBuffers[ci].(*DataSource.MemColumnBuffer).Append(res[k])
				ci++
			}
			ci = cii
		}
	}
	fmt.Println("======", len(columnBuffers))

	columnNames := []string{}
	for i := 0; i < len(self.SelectItems); i++ {
		item := self.SelectItems[i]
		columnNames = append(columnNames, item.GetNames()...)
	}

	return DataSource.NewDataSource([]string{self.Name}, columnNames, columnBuffers)
}

///////////////////
type PlanScanNode struct {
	Input *DataSource.DataSource
}

func NewPlanScanNodeFromDataSource(ctx *Context.Context, input *DataSource.DataSource) *PlanScanNode {
	res := &PlanScanNode{
		Input: input,
	}
	return res
}

func NewPlanScanNode(ctx *Context.Context, name string) *PlanScanNode {
	res := &PlanScanNode{
		Input: ctx.Tables[name],
	}
	return res
}

func (self *PlanScanNode) Execute() *DataSource.DataSource {
	return self.Input
}
