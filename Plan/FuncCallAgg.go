package Plan

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

func AggLocalFuncToAggGlobalFunc(f *GueryFunc) *GueryFunc {
	switch f.Name {
	case "COUNT":
		return NewCountGlobalFunc()
	case "SUM":
		return NewSumGlobalFunc()
	case "AVG":
		return NewAvgGlobalFunc()
	case "MIN":
		return NewMinGlobalFunc()
	case "MAX":
		return NewMaxGlobalFunc()
	}
	return nil
}

func NewCountGlobalFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "COUNTGLOBAL",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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
					key := row.GetKeyString()
					if _, ok := funcRes[key]; !ok {
						funcRes[key] = tmp

					} else {
						funcRes[key] = Type.OperatorFunc(funcRes[key], tmp, Type.PLUS)
					}
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewCountFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "COUNT",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				key := row.GetKeyString()
				if tmp != nil {
					if _, ok := funcRes[key]; !ok {
						funcRes[key] = int64(0)
					}
					funcRes[key] = funcRes[key].(int64) + 1
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewSumGlobalFunc() *GueryFunc {
	res := NewSumFunc()
	res.Name = "SUMGLOBAL"
	return res

}

func NewSumFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "SUM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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

				key := row.GetKeyString()
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = tmp
				} else {
					funcRes[key] = Type.OperatorFunc(funcRes[key], tmp, Type.PLUS)
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewAvgGlobalFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "AVGGLOBAL",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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
				t   = Expressions[0]
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
				if tmp, err = t.Result(rb); err != nil {
					break
				}
				key := row.GetKeyString()
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = tmp
				} else {
					var sumctmp, cntctmp float64
					fmt.Sscanf(tmp.(string), "%f:%f", &sumctmp, &cntctmp)
					var sumc, cntc float64
					fmt.Sscanf(funcRes[key].(string), "%f:%f", &sumc, &cntc)
					funcRes[key] = fmt.Sprintf("%v:%v", sumc+sumctmp, cntc+cntctmp)
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewAvgFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "AVG",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
		},

		GetType: func(md *Metadata.Metadata, es []*ExpressionNode) (Type.Type, error) {
			return Type.STRING, nil
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
				rb = Row.NewRowsGroup(input.Metadata)
				rb.Write(row)
				tmp, err = t.Result(rb)
				if err != nil {
					if err == io.EOF {
						err = nil
					}
					break
				}
				key := row.GetKeyString()
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = fmt.Sprintf("%v:%v", tmp, 1)

				} else {
					var sumc, cntc float64
					fmt.Sscanf(funcRes[key].(string), "%f:%f", &sumc, &cntc)
					sumc = sumc + Type.ToFloat64(tmp)
					cntc = cntc + float64(1)
					funcRes[key] = fmt.Sprintf("%v:%v", sumc, cntc)
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewMinGlobalFunc() *GueryFunc {
	res := NewMinFunc()
	res.Name = "MINGLOBAL"
	return res
}

func NewMinFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "MIN",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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

				key := row.GetKeyString()
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = tmp
				} else {
					if Type.GTFunc(funcRes[key], tmp).(bool) {
						funcRes[key] = tmp
					}
				}
			}
			return funcRes, err
		},
	}
	return res
}

func NewMaxGlobalFunc() *GueryFunc {
	res := NewMaxFunc()
	res.Name = "MAXGLOBAL"
	return res
}

func NewMaxFunc() *GueryFunc {
	var funcRes map[string]interface{}

	res := &GueryFunc{
		Name: "MAX",
		IsAggregate: func(ex []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
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

				key := row.GetKeyString()
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = tmp
				} else {
					if Type.LTFunc(funcRes[key], tmp).(bool) {
						funcRes[key] = tmp
					}
				}
			}
			return funcRes, err
		},
	}
	return res
}
