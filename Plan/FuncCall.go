package Plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type GueryFunc struct {
	Name        string
	Result      func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error)
	IsAggregate func(es []*ExpressionNode) bool
	GetType     func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error)
}

var Funcs map[string]*GueryFunc

func init() {
	Funcs = map[string]*GueryFunc{
		//aggregate functions
		"SUM": NewSumFunc(),
		"AVG": NewAvgFunc(),
		"MAX": NewMaxFunc(),
		"MIN": NewMinFunc(),

		//math functions
		"ABS":    NewAbsFunc(),
		"SQRT":   NewSqrtFunc(),
		"POW":    NewPowFunc(),
		"RAND":   NewRandomFunc(),
		"RANDOM": NewRandomFunc(),

		"LOG":   NewLogFunc(),
		"LOG10": NewLog10Func(),
		"LOG2":  NewLog2Func(),
		"LN":    NewLnFunc(),

		"FLOOR":   NewFloorFunc(),
		"CEIL":    NewCeilFunc(),
		"CEILING": NewCeilFunc(),
		"ROUND":   NewRoundFunc(),

		"SIN":  NewSinFunc(),
		"COS":  NewCosFunc(),
		"TAN":  NewTanFunc(),
		"ASIN": NewASinFunc(),
		"ACOS": NewACosFunc(),
		"ATAN": NewATanFunc(),

		"SINH":  NewSinhFunc(),
		"COSH":  NewCoshFunc(),
		"TANH":  NewTanhFunc(),
		"ASINH": NewASinhFunc(),
		"ACOSH": NewACoshFunc(),
		"ATANH": NewATanhFunc(),

		"E":  NewEFunc(),
		"PI": NewPiFunc(),

		//string functions
		"LENGTH":  NewLengthFunc(),
		"LOWER":   NewLowerFunc(),
		"UPPER":   NewUpperFunc(),
		"CONCAT":  NewConcatFunc(),
		"REVERSE": NewReverseFunc(),
		"SUBSTR":  NewSubstrFunc(),
		"REPLACE": NewReplaceFunc(),
	}
}

////////////////////////

type FuncCallNode struct {
	FuncName    string
	Expressions []*ExpressionNode
}

func NewFuncCallNode(name string, expressions []parser.IExpressionContext) *FuncCallNode {
	res := &FuncCallNode{
		FuncName:    strings.ToUpper(name),
		Expressions: make([]*ExpressionNode, len(expressions)),
	}
	for i := 0; i < len(expressions); i++ {
		res.Expressions[i] = NewExpressionNode(expressions[i])
	}
	return res
}

func (self *FuncCallNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	if fun, ok := Funcs[self.FuncName]; ok {
		return fun.Result(input, self.Expressions)
	}
	return nil, fmt.Errorf("Unkown function %v", self.FuncName)
}

func (self *FuncCallNode) GetType(md *Util.Metadata) (Util.Type, error) {
	if fun, ok := Funcs[self.FuncName]; ok {
		return fun.GetType(md, self.Expressions)
	}
	return Util.UNKNOWNTYPE, fmt.Errorf("Unkown function %v", self.FuncName)
}

func (self *FuncCallNode) GetColumns() ([]string, error) {
	res, resmp := []string{}, map[string]int{}
	for _, e := range self.Expressions {
		cs, err := e.GetColumns()
		if err != nil {
			return res, err
		}
		for _, c := range cs {
			resmp[c] = 1
		}
	}
	for c, _ := range resmp {
		res = append(res, c)
	}
	return res, nil
}

func (self *FuncCallNode) IsAggregate() bool {
	if fun, ok := Funcs[self.FuncName]; ok {
		return fun.IsAggregate(self.Expressions)
	}
	return false
}
