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
	NOT
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

func Less(va interface{}, vb interface{}) bool {
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
