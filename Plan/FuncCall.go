package Plan

import (
	"github.com/xitongsys/guery/Common"
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type FuncCallNode struct {
	FuncName    string
	Expressions []*ExpressionNode
}

func NewFuncCallNode(ctx *Context.Context, name string, expressions []parser.IExpressionContext) *FuncCallNode {
	res := &FuncCallNode{
		FuncName:    name,
		Expressions: make([]*ExpressionNode, len(expressions)),
	}
	for i := 0; i < len(expressions); i++ {
		res.Expressions[i] = NewExpressionNode(ctx, expressions[i])
	}
	return res
}

func (self *FuncCallNode) Result(input *DataSource.DataSource) interface{} {
	switch self.FuncName {
	case "SUM":
		return SUM(input, self.Expressions[0])
	case "MIN":
		return MIN(input, self.Expressions[0])
	case "MAX":
		return MAX(input, self.Expressions[0])
	case "ABS":
		return ABS(input)
	}
	return nil
}

func (self *FuncCallNode) IsAggregate() bool {
	switch self.FuncName {
	case "SUM":
		return true
	case "MIN":
		return true
	case "MAX":
		return true
	case "ABS":
		return false
	}
	return false
}

func SUM(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			res = Common.Arithmetic(res, tmp, Common.PLUS)
		}
		input.Next()
	}
	return res
}

func MIN(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			if Common.Cmp(res, tmp) > 0 {
				res = tmp
			}
		}
		input.Next()
	}
	return res
}

func MAX(input *DataSource.DataSource, t *ExpressionNode) interface{} {
	var res interface{}
	for !input.IsEnd() {
		tmp := t.Result(input)
		if res == nil {
			res = tmp
		} else {
			if Common.Cmp(res, tmp) < 0 {
				res = tmp
			}
		}
		input.Next()
	}
	return res
}

func ABS(input *DataSource.DataSource) interface{} {
	return nil
}
