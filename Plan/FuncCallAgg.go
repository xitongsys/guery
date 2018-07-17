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
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if esi, err = t.Result(rb); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
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
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if esi, err = t.Result(input); err != nil {
				break
			}
			es := esi.([]interface{})
			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = int64(0)
				}
				funcRes[key] = funcRes[key].(int64) + 1
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
	var valType Type.Type

	res := &GueryFunc{
		Name: "SUM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return true
		},

		Init: func() {
			funcRes = make(map[string]interface{})
			valType = Type.UNKNOWNTYPE
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
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if esi, err = t.Result(input); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = es[i]
				} else {
					funcRes[key] = Type.OperatorFunc(funcRes[key], es[i], Type.PLUS)
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
				esi interface{}
				t   = Expressions[0]
			)

			if esi, err = t.Result(input); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = es[i]
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
			return Type.FLOAT64, nil
		},

		Result: func(input *Row.RowsGroup, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in AVG")
			}
			var (
				err error
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if esi, err = t.Result(rb); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
				if _, ok := funcRes[key]; !ok {
					funcRes[key] = fmt.Sprintf("%v:%v", es[i], 1)
				} else {
					var sumc, cntc float64
					fmt.Sscanf(funcRes[key].(string), "%f:%f", &sumc, &cntc)
					sumc = sumc + Type.ToFloat64(es[i])
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
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if esi, err = t.Result(rb); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
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
				esi interface{}
				t   *ExpressionNode = Expressions[0]
			)
			if esi, err = t.Result(rb); err != nil {
				break
			}
			es := esi.([]interface{})

			for i := 0; i < len(es); i++ {
				key := input.GetKeyString(i)
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
