package plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type GueryFunc struct {
	Name        string
	Result      func(input *row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error)
	IsAggregate func(es []*ExpressionNode) bool
	GetType     func(md *metadata.Metadata, es []*ExpressionNode) (gtype.Type, error)
	Init        func()
}

var Funcs map[string](func() *GueryFunc)

func init() {
	Funcs = map[string](func() *GueryFunc){
		//aggregate functions
		"SUM":         NewSumFunc,
		"SUMGLOBAL":   NewSumGlobalFunc,
		"AVG":         NewAvgFunc,
		"AVGGLOBAL":   NewAvgGlobalFunc,
		"MAX":         NewMaxFunc,
		"MAXGLOBAL":   NewMaxGlobalFunc,
		"MIN":         NewMinFunc,
		"MINGLOBAL":   NewMinGlobalFunc,
		"COUNT":       NewCountFunc,
		"COUNTGLOBAL": NewCountGlobalFunc,

		//math functions
		"ABS":    NewAbsFunc,
		"SQRT":   NewSqrtFunc,
		"POW":    NewPowFunc,
		"RAND":   NewRandomFunc,
		"RANDOM": NewRandomFunc,

		"LOG":   NewLogFunc,
		"LOG10": NewLog10Func,
		"LOG2":  NewLog2Func,
		"LN":    NewLnFunc,

		"FLOOR":   NewFloorFunc,
		"CEIL":    NewCeilFunc,
		"CEILING": NewCeilFunc,
		"ROUND":   NewRoundFunc,

		"SIN":  NewSinFunc,
		"COS":  NewCosFunc,
		"TAN":  NewTanFunc,
		"ASIN": NewASinFunc,
		"ACOS": NewACosFunc,
		"ATAN": NewATanFunc,

		"SINH":  NewSinhFunc,
		"COSH":  NewCoshFunc,
		"TANH":  NewTanhFunc,
		"ASINH": NewASinhFunc,
		"ACOSH": NewACoshFunc,
		"ATANH": NewATanhFunc,

		"E":  NewEFunc,
		"PI": NewPiFunc,

		//string functions
		"LENGTH":  NewLengthFunc,
		"LOWER":   NewLowerFunc,
		"UPPER":   NewUpperFunc,
		"CONCAT":  NewConcatFunc,
		"REVERSE": NewReverseFunc,
		"SUBSTR":  NewSubstrFunc,
		"REPLACE": NewReplaceFunc,

		//time functions
		"NOW":    NewNowFunc,
		"DAY":    NewDayFunc,
		"MONTH":  NewMonthFunc,
		"YEAR":   NewYearFunc,
		"HOUR":   NewHourFunc,
		"MINUTE": NewMinuteFunc,
		"SECOND": NewSecondFunc,
	}
}

////////////////////////

type FuncCallNode struct {
	FuncName    string
	ResColName  string //used in ExtractAggFunc
	Func        *GueryFunc
	Expressions []*ExpressionNode
}

func NewFuncCallNode(runtime *config.ConfigRuntime, name string, expressions []parser.IExpressionContext) *FuncCallNode {
	name = strings.ToUpper(name)
	res := &FuncCallNode{
		FuncName:    name,
		Expressions: make([]*ExpressionNode, len(expressions)),
	}
	for i := 0; i < len(expressions); i++ {
		res.Expressions[i] = NewExpressionNode(runtime, expressions[i])
	}

	return res
}

func (self *FuncCallNode) Init(md *metadata.Metadata) error {
	for _, e := range self.Expressions {
		if err := e.Init(md); err != nil {
			return err
		}
	}

	if self.Func == nil {
		if f, ok := Funcs[self.FuncName]; ok {
			self.Func = f()
		} else {
			return fmt.Errorf("Unknown function %v", self.FuncName)
		}
	}
	self.Func.Init()
	return nil
}

func (self *FuncCallNode) Result(input *row.RowsGroup) (interface{}, error) {
	if self.Func != nil {
		return self.Func.Result(input, self.Expressions)
	}
	return nil, fmt.Errorf("Unkown function %v", self.FuncName)
}

func (self *FuncCallNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	if fun, ok := Funcs[self.FuncName]; ok {
		return fun().GetType(md, self.Expressions)
	}
	return gtype.UNKNOWNTYPE, fmt.Errorf("Unkown function %v", self.FuncName)
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
		return fun().IsAggregate(self.Expressions)
	}
	return false
}
