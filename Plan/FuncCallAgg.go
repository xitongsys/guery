package Plan

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
	"github.com/xitongsys/guery/Type"
)

func NewCountFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "COUNT",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT64, nil
		},

		Result: func(input *Split.Split, index int, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in SUM")
			}
			var (
				err error
				res int64
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			for i := index; i < input.GetRowsNumber(); i++ {
				tmp, err = t.Result(input, i)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				if tmp != nil {
					res = res + 1
				}
			}
			return res, err
		},
	}
	return res
}

func NewSumFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SUM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in SUM")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Split.Split, index int, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in SUM")
			}
			var (
				err      error
				res, tmp interface{}
				t        *ExpressionNode = Expressions[0]
			)

			for i := index; i < input.GetRowsNumber(); i++ {
				tmp, err = t.Result(input, i)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if res == nil {
					res = tmp
				} else {
					res = Type.OperatorFunc(res, tmp, Type.PLUS)
				}
			}
			return res, err
		},
	}
	return res
}

func NewAvgFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "AVG",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.FLOAT64, nil
		},

		Result: func(input *Split.Split, index int, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in AVG")
			}
			var (
				err      error
				res, tmp interface{}
				cnt      float64
				t        *ExpressionNode = Expressions[0]
			)

			for i := index; i < input.GetRowsNumber(); i++ {
				cnt++
				tmp, err = t.Result(input, i)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if res == nil {
					res = tmp
				} else {
					res = Type.OperatorFunc(res, tmp, Type.PLUS)
				}
			}
			if cnt > 0 {
				var cnti interface{} = cnt
				res = Type.OperatorFunc(res, cnti, Type.SLASH)
			}
			return res, err
		},
	}
	return res
}

func NewMinFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "MIN",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in MIN")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Split.Split, index int, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MIN")
			}
			var (
				err      error
				res, tmp interface{}
				t        *ExpressionNode = Expressions[0]
			)

			for i := index; i < input.GetRowsNumber(); i++ {
				tmp, err = t.Result(input, i)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if res == nil {
					res = tmp
				} else {
					if Type.GTFunc(res, tmp).(bool) {
						res = tmp
					}
				}
			}
			return res, err
		},
	}
	return res
}

func NewMaxFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "MAX",
		IsAggregate: func(ex []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			if len(es) < 1 {
				return Type.UNKNOWNTYPE, fmt.Errorf("not enough parameters in MAX")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Split.Split, index int, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MAX")
			}
			var (
				err      error
				res, tmp interface{}
				t        *ExpressionNode = Expressions[0]
			)

			for i := index; i < input.GetRowsNumber(); i++ {
				tmp, err = t.Result(input, i)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}

				if res == nil {
					res = tmp
				} else {
					if Type.LTFunc(res, tmp).(bool) {
						res = tmp
					}
				}
			}
			return res, err
		},
	}
	return res
}
