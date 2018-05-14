package Util

import (
	"fmt"
)

//////////////////
type Type int32

const (
	UNKNOWNTYPE Type = iota
	BOOL
	INT
	DOUBLE
	STRING
)

///////////////////
type OrderType int32

const (
	UNKNOWNORDERTYPE OrderType = iota
	ASC
	DESC
	FIRST
	LAST
)

////////////////
type FuncType int32

const (
	UNKNOWNFUNCTYPE FuncType = iota
	AGGREGATE
	NORMAL
)

/////////////
type QuantifierType int32

const (
	UNKNOWNQUANTIFIERTYPE QuantifierType = iota
	ALL
	DISTINCT
	SOME
	ANY
)

func CheckType(ta, tb Type, op Operator) (Type, error) {
	if ta != tb || ta == UNKNOWNTYPE {
		return UNKNOWNTYPE, fmt.Errorf("type not match")
	}
	return ta, nil
}

/////////////////////////
func IsInt(va interface{}) bool {
	_, ok := va.(int64)
	return ok
}

func ToInt(va interface{}) (int64, bool) {
	if v, ok := va.(int64); ok {
		return v, true
	}
	if v, ok := va.(float64); ok {
		return int64(v), true
	}
	return 0, false
}

func IsDouble(va interface{}) bool {
	_, ok := va.(float64)
	return ok
}

func ToDouble(va interface{}) (float64, bool) {
	if v, ok := va.(float64); ok {
		return v, true
	}
	if v, ok := va.(int64); ok {
		return float64(v), true
	}
	return 0, false
}

func Less(va interface{}, vb interface{}) bool {
	if va == nil || vb != nil {
		return true
	} else if va != nil || vb == nil {
		return false
	} else if va == nil && vb == nil {
		return false
	} else {
		if IsDouble(va) || IsDouble(vb) {
			a, oka := ToDouble(va)
			b, okb := ToDouble(vb)
			if !oka || !okb {
				return false
			}
			return a < b
		}
		if IsInt(va) || IsInt(vb) {
			a, oka := ToInt(va)
			b, okb := ToInt(vb)
			if !oka || !okb {
				return false
			}
			return a < b
		}
		a, oka := va.(string)
		b, okb := vb.(string)
		if !oka || !okb {
			return false
		}
		return a < b
	}
}
