package Plan

import (
	"fmt"
	"strings"

	"github.com/xitongsys/guery/Util"
)

func NewLengthFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "LENGTH",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.INT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in LENGTH")
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
			case Util.STRING:
				return int64(len(tmp.(string))), nil

			default:
				return nil, fmt.Errorf("type cann't use LENGTH function")
			}
		},
	}
	return res
}

func NewLowerFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "LOWER",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.STRING, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in LOWER")
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
			case Util.STRING:
				return strings.ToLower(tmp.(string)), nil

			default:
				return nil, fmt.Errorf("type cann't use LOWER function")
			}
		},
	}
	return res
}

func NewUpperFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "UPPER",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.STRING, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in UPPER")
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
			case Util.STRING:
				return strings.ToUpper(tmp.(string)), nil

			default:
				return nil, fmt.Errorf("type cann't use UPPER function")
			}
		},
	}
	return res
}

func NewConcatFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "CONCAT",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 2 {
				return false
			}
			return es[0].IsAggregate() || es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.STRING, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 2 {
				return nil, fmt.Errorf("not enough parameters in CONCAT")
			}
			var (
				err        error
				tmp1, tmp2 interface{}
				t1         *ExpressionNode = Expressions[0]
				t2         *ExpressionNode = Expressions[1]
			)

			input.Reset()
			if tmp1, err = t1.Result(input); err != nil {
				return nil, err
			}

			input.Reset()
			if tmp2, err = t2.Result(input); err != nil {
				return nil, err
			}

			if Util.TypeOf(tmp1) != Util.STRING || Util.TypeOf(tmp2) != Util.STRING {
				return nil, fmt.Errorf("type error in CONCAT")
			}

			return tmp1.(string) + tmp2.(string), nil
		},
	}
	return res
}
