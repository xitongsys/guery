package Plan

import (
	"fmt"
	"math"

	"github.com/xitongsys/guery/Util"
)

func NewAbsFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ABS",
		IsAggregate: func(es []*ExpressionNode) bool {
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
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
		},
	}
	return res
}

func NewSqrtFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SQRT",
		IsAggregate: func(es []*ExpressionNode) bool {
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return nil, fmt.Errorf("type cann't use SQRT function")

			default:
				return math.Sqrt(Util.ToFloat64(tmp)), nil
			}
		},
	}
	return res
}
