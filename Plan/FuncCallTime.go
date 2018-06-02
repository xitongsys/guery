package Plan

import (
	"fmt"
	"time"

	"github.com/xitongsys/guery/Util"
)

func NewNowFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "NOW",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.TIMESTAMP, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			return time.Now(), nil
		},
	}
	return res
}

func NewDayFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "DAY",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.INT32, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in DAY")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			input.Reset()
			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.TIMESTAMP:
				return int32(tmp.(time.Time).Day()), nil
			default:
				return nil, fmt.Errorf("type cann't use DAY function")
			}
		},
	}
	return res
}

func NewMonthFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "MONTH",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.INT32, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MONTH")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			input.Reset()
			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.TIMESTAMP:
				return int32(tmp.(time.Time).Month()), nil
			default:
				return nil, fmt.Errorf("type cann't use MONTH function")
			}
		},
	}
	return res
}

func NewYearFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "YEAR",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.INT32, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in YEAR")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			input.Reset()
			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.TIMESTAMP:
				return int32(tmp.(time.Time).Year()), nil
			default:
				return nil, fmt.Errorf("type cann't use YEAR function")
			}
		},
	}
	return res
}
