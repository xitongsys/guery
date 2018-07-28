package plan

import (
	"fmt"
	"time"

	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/row"
)

func NewNowFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "NOW",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.TIMESTAMP, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
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

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
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

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
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

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
				return int32(tmp.(time.Time).Year()), nil
			default:
				return nil, fmt.Errorf("type cann't use YEAR function")
			}
		},
	}
	return res
}

func NewHourFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "HOUR",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in HOUR")
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
				return int32(tmp.(time.Time).Hour()), nil
			default:
				return nil, fmt.Errorf("type cann't use HOUR function")
			}
		},
	}
	return res
}

func NewMinuteFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "MINUTE",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in MINUTE")
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
				return int32(tmp.(time.Time).Minute()), nil
			default:
				return nil, fmt.Errorf("type cann't use MINUE function")
			}
		},
	}
	return res
}

func NewSecondFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SECOND",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.INT32, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in SECOND")
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

			switch Type.TypeOf(tmp) {
			case Type.TIMESTAMP:
				return int32(tmp.(time.Time).Second()), nil
			default:
				return nil, fmt.Errorf("type cann't use SECOND function")
			}
		},
	}
	return res
}
