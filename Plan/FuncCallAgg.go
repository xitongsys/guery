package Plan

import (
	"io"

	"github.com/xitongsys/guery/Util"
)

func NewSumFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SUM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err      error
				res, tmp interface{}
				rb       *Util.RowsBuffer
				row      *Util.Row
				t        *ExpressionNode = Expressions[0]
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

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err      error
				res, tmp interface{}
				rb       *Util.RowsBuffer
				row      *Util.Row
				cnt      float64
				t        *ExpressionNode = Expressions[0]
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
			if cnt > 0 {
				var cnti interface{} = cnt
				res = Util.OperatorFunc(res, cnti, Util.SLASH)
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

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err      error
				res, tmp interface{}
				rb       *Util.RowsBuffer
				row      *Util.Row
				t        *ExpressionNode = Expressions[0]
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

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			var (
				err      error
				res, tmp interface{}
				rb       *Util.RowsBuffer
				row      *Util.Row
				t        *ExpressionNode = Expressions[0]
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
		},
	}
	return res
}
