package Util

import (
	"fmt"
	"time"
)

///////////////////////
type Operator int32

const (
	_ Operator = iota
	AND
	OR
	ASTERISK
	SLASH
	PERCENT
	PLUS
	MINUS
	EQ
	NEQ
	LT
	LTE
	GT
	GTE
)

func NewOperatorFromString(name string) *Operator {
	var res Operator
	switch name {
	case "AND":
		res = AND
	case "OR":
		res = OR
	case "NOT":
		res = NOT
	case "*":
		res = ASTERISK
	case "/":
		res = SLASH
	case "%":
		res = PERCENT
	case "+":
		res = PLUS
	case "-":
		res = MINUS
	case ">":
		res = GT
	case "<":
		res = LT
	case ">=":
		res = GTE
	case "<=":
		res = LTE
	case "=":
		res = EQ
	case "<>":
		res = NEQ
	}
	return &res
}

//+
func PLUSFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			return ToInt32(va) + ToInt32(vb)
		case INT32:
			return va.(int32) + vb.(int32)
		case INT64:
			return va.(int64) + vb.(int64)
		case FLOAT32:
			return va.(float32) + vb.(float32)
		case FLOAT64:
			return va.(float64) + vb.(float64)
		case STRING:
			return va.(string) + vb.(string)
		case TIMESTAMP:
			return time.Unix(va.(time.Time).Unix()+vb.(time.Time).Unix(), 0).Unix()
		}
	}
	return nil
}

//-
func MINUSFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			return ToInt32(va) - ToInt32(vb)
		case INT32:
			return va.(int32) - vb.(int32)
		case INT64:
			return va.(int64) - vb.(int64)
		case FLOAT32:
			return va.(float32) - vb.(float32)
		case FLOAT64:
			return va.(float64) - vb.(float64)
		case STRING:
			return nil
		case TIMESTAMP:
			return time.Unix(va.(time.Time).Unix()-vb.(time.Time).Unix(), 0).Unix()
		}
	}
	return nil
}

//*
func ASTERISKFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			return ToInt32(va) * ToInt32(vb)
		case INT32:
			return va.(int32) * vb.(int32)
		case INT64:
			return va.(int64) * vb.(int64)
		case FLOAT32:
			return va.(float32) * vb.(float32)
		case FLOAT64:
			return va.(float64) * vb.(float64)
		case STRING:
			return nil
		case TIMESTAMP:
			return nil
		}
	}
	return nil
}

///
func SLASHFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			return ToInt32(va) / ToInt32(vb)
		case INT32:
			return va.(int32) / vb.(int32)
		case INT64:
			return va.(int64) / vb.(int64)
		case FLOAT32:
			return va.(float32) / vb.(float32)
		case FLOAT64:
			return va.(float64) / vb.(float64)
		case STRING:
			return nil
		case TIMESTAMP:
			return nil
		}
	}
	return nil
}

//%
func PERCENTFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			return ToInt32(va) % ToInt32(vb)
		case INT32:
			return va.(int32) % vb.(int32)
		case INT64:
			return va.(int64) % vb.(int64)
		case FLOAT32:
			return nil
		case FLOAT64:
			return nil
		case STRING:
			return nil
		case TIMESTAMP:
			return nil
		}
	}
	return nil
}

//AND
func ANDFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToBool(va), ToBool(vb)
		return va && vb
	}
}

//OR
func ORFunc(va interface{}, vb interface{}) interface{} {
	if va == nil || vb == nil {
		return nil
	} else {
		va, vb = ToBool(va), ToBool(vb)
		return va || vb
	}
}

//==
func EQFunc(va interface{}, vb interface{}) interface{} {
	return va == vb
}

//<>
func NEQFunc(va interface{}, vb interface{}) interface{} {
	return !EQ(va, vb)
}

//<
func LTFunc(va interface{}, vb interface{}) interface{} {
	if va == nil && vb != nil {
		return true
	} else if va != nil && vb == nil {
		return false
	} else if va == nil && vb == nil {
		return false
	} else if va == vb {
		return false
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			a, b := va.(bool), vb.(bool)
			if !a && b {
				return true
			}
			return false
		case INT32:
			return va.(int32) < vb.(int32)
		case INT64:
			return va.(int64) < vb.(int64)
		case FLOAT32:
			return va.(float32) < vb.(float32)
		case FLOAT64:
			return va.(float64) < vb.(float64)
		case STRING:
			return va.(string) < vb.(string)
		case TIMESTAMP:
			a, b := va.(time.Time), vb.(time.Time)
			return a.Before(b)
		}
	}
	return false
}

//<=
func LTEFunc(va interface{}, vb interface{}) interface{} {
	if va == nil && vb != nil {
		return true
	} else if va != nil && vb == nil {
		return false
	} else if va == nil && vb == nil {
		return false
	} else if va == vb {
		return false
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			a, b := va.(bool), vb.(bool)
			if (!a && b) || (a == b) {
				return true
			}
			return false
		case INT32:
			return va.(int32) <= vb.(int32)
		case INT64:
			return va.(int64) <= vb.(int64)
		case FLOAT32:
			return va.(float32) <= vb.(float32)
		case FLOAT64:
			return va.(float64) <= vb.(float64)
		case STRING:
			return va.(string) <= vb.(string)
		case TIMESTAMP:
			a, b := va.(time.Time), vb.(time.Time)
			return !b.Before(a)
		}
	}
	return false
}

//>
func GTFunc(va interface{}, vb interface{}) interface{} {
	if va == nil && vb != nil {
		return true
	} else if va != nil && vb == nil {
		return false
	} else if va == nil && vb == nil {
		return false
	} else if va == vb {
		return false
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			a, b := va.(bool), vb.(bool)
			if a && !b {
				return true
			}
			return false
		case INT32:
			return va.(int32) > vb.(int32)
		case INT64:
			return va.(int64) > vb.(int64)
		case FLOAT32:
			return va.(float32) > vb.(float32)
		case FLOAT64:
			return va.(float64) > vb.(float64)
		case STRING:
			return va.(string) > vb.(string)
		case TIMESTAMP:
			a, b := va.(time.Time), vb.(time.Time)
			return a.After(b)
		}
	}
	return false
}

//>=
func GTEFunc(va interface{}, vb interface{}) interface{} {
	if va == nil && vb != nil {
		return true
	} else if va != nil && vb == nil {
		return false
	} else if va == nil && vb == nil {
		return false
	} else if va == vb {
		return false
	} else {
		va, vb = ToSameType(va, vb)
		t := TypeOf(va)
		switch t {
		case BOOL:
			a, b := va.(bool), vb.(bool)
			if (a && !b) || (a == b) {
				return true
			}
			return false
		case INT32:
			return va.(int32) >= vb.(int32)
		case INT64:
			return va.(int64) >= vb.(int64)
		case FLOAT32:
			return va.(float32) >= vb.(float32)
		case FLOAT64:
			return va.(float64) >= vb.(float64)
		case STRING:
			return va.(string) >= vb.(string)
		case TIMESTAMP:
			a, b := va.(time.Time), vb.(time.Time)
			return !b.After(a)
		}
	}
	return false
}

func OperatorFunc(va interface{}, vb interface{}, op Operator) interface{} {
	switch op {
	case AND:
		return ANDFunc(va, vb)
	case OR:
		return ORFunc(va, vb)
	case ASTERISK:
		return ASTERISKFunc(va, vb)
	case SLASH:
		return SLASHFunc(va, vb)
	case PERCENT:
		return PERCENTFunc(va, vb)
	case PLUS:
		return PLUSFunc(va, vb)
	case MINUS:
		return MINUSFunc(va, vb)
	case EQ:
		return EQFunc(va, vb)
	case NEQ:
		return NEQFunc(va, vb)
	case LT:
		return LTFunc(va, vb)
	case LTE:
		return LTEFunc(va, vb)
	case GT:
		return GTFunc(va, vb)
	case GTE:
		return GTEFunc(va, vb)
	}
	return nil
}
