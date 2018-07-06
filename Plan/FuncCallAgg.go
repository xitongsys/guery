package Plan

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

func NewCountFunc() *GueryFunc {
	var funcRes int64

	res := &GueryFunc{
		Name: "COUNT",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = 0
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT64, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in SUM")
			}
			var (
				err error
				tmp interface{}
				rb  *Row.RowsGroup
				row *Row.Row
				t   *ExpressionNode = Expressions[0]
			)

			for {
				row, err = input.Read()
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				if tmp != nil {
					funcRes = funcRes + 1
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewSumFunc() *GueryFunc {
	var funcRes interface{}

	res := &GueryFunc{
		Name: "SUM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = nil
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in SUM")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in SUM")
			}
			var (
				err error
				tmp interface{}
				rb  *Row.RowsGroup
				row *Row.Row
				t   *ExpressionNode = Expressions[0]
			)

			for {
				row, err = input.Read()
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if funcRes == nil {
					funcRes = tmp
				} else {
					funcRes = Type.OperatorFunc(funcRes, tmp, Type.PLUS)
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewAvgFunc() *GueryFunc {
	var funcRes interface{}
	var cnt float64

	res := &GueryFunc{
		Name: "AVG",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = nil
			cnt = float64(0)
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.FLOAT64, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in AVG")
			}
			var (
				err error
				tmp interface{}
				rb  *Row.RowsGroup
				row *Row.Row
				t   *ExpressionNode = Expressions[0]
			)

			for {
				row, err = input.Read()
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				cnt++
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if funcRes == nil {
					funcRes = tmp
				} else {
					funcRes = Type.OperatorFunc(funcRes, tmp, Type.PLUS)
				}
			}
			if cnt > 0 {
				var cnti interface{} = cnt
				funcRes = Type.OperatorFunc(funcRes, cnti, Type.SLASH)
			}
			return funcRes, err
		},
	}
	return res
}

func NewMinFunc() *GueryFunc {
	var funcRes interface{}

	res := &GueryFunc{
		Name: "MIN",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = nil
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in MIN")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MIN")
			}
			var (
				err error
				tmp interface{}
				rb  *Row.RowsGroup
				row *Row.Row
				t   *ExpressionNode = Expressions[0]
			)

			for {
				row, err = input.Read()
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if funcRes == nil {
					funcRes = tmp
				} else {
					if Type.GTFunc(funcRes, tmp).(bool) {
						funcRes = tmp
					}
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewMaxFunc() *GueryFunc {
	var funcRes interface{}

	res := &GueryFunc{
		Name: "MAX",
		IsAggregate: func(ex []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = nil
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in MAX")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MAX")
			}
			var (
				err error
				tmp interface{}
				rb  *Row.RowsGroup
				row *Row.Row
				t   *ExpressionNode = Expressions[0]
			)

			for {
				row, err = input.Read()
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if funcRes == nil {
					funcRes = tmp
				} else {
					if Type.LTFunc(funcRes, tmp).(bool) {
						funcRes = tmp
					}
				}
			}
			return funcRes, err
		},
	}
	return res
}
