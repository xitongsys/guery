package gtype

import (
	"fmt"
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
