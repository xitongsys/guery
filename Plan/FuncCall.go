package Plan

import (
	"fmt"
	"io"
	"strings"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

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
	switch self.FuncName {
	case "SUM":
		return SUM(input, self.Expressions[0])
	case "MIN":
		return MIN(input, self.Expressions[0])
	case "MAX":
		return MAX(input, self.Expressions[0])
	case "ABS":
		return ABS(input, self.Expressions[0])
	}
	return nil, fmt.Errorf("Unkown function %v", self.FuncName)
}

func (self *FuncCallNode) GetType(md *Util.Metadata) (Util.Type, error) {
	switch self.FuncName {
	case "SUM":
		return SUMType(md, self.Expressions[0])
	case "MIN":
		return MINType(md, self.Expressions[0])
	case "MAX":
		return MAXType(md, self.Expressions[0])
	case "ABS":
		return ABSType(md, self.Expressions[0])
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
	switch self.FuncName {
	case "SUM":
		return true
	case "MIN":
		return true
	case "MAX":
		return true
	case "ABS":
		return self.Expressions[0].IsAggregate()
	}
	return false
}

func SUMType(md *Util.Metadata, t *ExpressionNode) (Util.Type, error) {
	return t.GetType(md)
}

func SUM(input *Util.RowsBuffer, t *ExpressionNode) (interface{}, error) {
	var (
		err      error
		res, tmp interface{}
		rb       *Util.RowsBuffer
		row      *Util.Row
	)

	for {
		row, err = input.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		rb = Util.NewRowsBuffer(input.Metadata)
		rb.Write(row)
		tmp, err = t.Result(rb)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if res == nil {
			res = tmp
		} else {
			res = Util.OperatorFunc(res, tmp, Util.PLUS)
		}
	}

	return res, err
}

func MINType(md *Util.Metadata, t *ExpressionNode) (Util.Type, error) {
	return t.GetType(md)
}

func MIN(input *Util.RowsBuffer, t *ExpressionNode) (interface{}, error) {
	var (
		err      error
		res, tmp interface{}
		rb       *Util.RowsBuffer
		row      *Util.Row
	)

	for {
		row, err = input.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		rb = Util.NewRowsBuffer(input.Metadata)
		rb.Write(row)
		tmp, err = t.Result(rb)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if res == nil {
			res = tmp
		} else {
			if Util.GTFunc(res, tmp).(bool) {
				res = tmp
			}
		}
	}
	return res, err
}

func MAXType(md *Util.Metadata, t *ExpressionNode) (Util.Type, error) {
	return t.GetType(md)
}

func MAX(input *Util.RowsBuffer, t *ExpressionNode) (interface{}, error) {
	var (
		err      error
		res, tmp interface{}
		rb       *Util.RowsBuffer
		row      *Util.Row
	)

	for {
		row, err = input.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		rb = Util.NewRowsBuffer(input.Metadata)
		rb.Write(row)
		tmp, err = t.Result(rb)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if res == nil {
			res = tmp
		} else {
			if Util.LTFunc(res, tmp).(bool) {
				res = tmp
			}
		}
	}
	return res, err
}

func ABSType(md *Util.Metadata, t *ExpressionNode) (Util.Type, error) {
	return t.GetType(md)
}
func ABS(input *Util.RowsBuffer, t *ExpressionNode) (interface{}, error) {
	var (
		err error
		tmp interface{}
	)

	if tmp, err = t.Result(input); err != nil {
		return nil, err
	}

	switch Util.TypeOf(tmp) {
	case Util.STRING, Util.BOOL, Util.TIMESTAMP:
		return nil, fmt.Errorf("type cann't use ABS function")
	case Util.FLOAT64:
		v := tmp.(float64)
		if v < 0 {
			v *= -1
		}
		return v, nil
	case Util.FLOAT32:
		v := tmp.(float32)
		if v < 0 {
			v *= -1
		}
		return v, nil
	case Util.INT64:
		v := tmp.(int64)
		if v < 0 {
			v *= -1
		}
		return v, nil
	case Util.INT32:
		v := tmp.(int32)
		if v < 0 {
			v *= -1
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unknown type")
	}
}
