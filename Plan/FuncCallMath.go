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
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			if len(es) < 1 {
				return Util.UNKNOWNTYPE, fmt.Errorf("not enough parameters in Abs")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in Abs")
			}
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
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in SQRT")
			}
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

func NewPowFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "POW",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 2 {
				return float64(0), fmt.Errorf("not enough parameters in POW")
			}
			var (
				err        error
				tmp1, tmp2 interface{}
				t1         *ExpressionNode = Expressions[0]
				t2         *ExpressionNode = Expressions[1]
			)

			if tmp1, err = t1.Result(input); err != nil {
				return float64(0), err
			}
			if tmp2, err = t2.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp1) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use SQRT function")

			default:
				switch Util.TypeOf(tmp2) {
				case Util.STRING, Util.BOOL, Util.TIMESTAMP:
					return float64(0), fmt.Errorf("type cann't use SQRT function")
				}
				v1, v2 := Util.ToFloat64(tmp1), Util.ToFloat64(tmp2)
				return math.Pow(v1, v2), nil
			}
		},
	}
	return res
}
