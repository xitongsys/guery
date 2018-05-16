package Util

import (
	"fmt"
	"reflect"
)

func Arithmetic(leftVal interface{}, rightVal interface{}, op Operator) interface{} {
	var res interface{}
	lT, rT := reflect.TypeOf(leftVal).Kind(), reflect.TypeOf(rightVal).Kind()
	if lT == reflect.Float64 || rT == reflect.Float64 {
		lv, rv := Val2Float(leftVal), Val2Float(rightVal)
		switch op {
		case ASTERISK:
			res = lv * rv
		case SLASH:
			res = lv / rv
		case PERCENT:
			res = nil
		case PLUS:
			res = lv + rv
		case MINUS:
			res = lv - rv
		}

	} else {
		lv, rv := Val2Int(leftVal), Val2Int(rightVal)
		switch op {
		case ASTERISK:
			res = lv * rv
		case SLASH:
			res = lv / rv
		case PERCENT:
			res = lv % rv
		case PLUS:
			res = lv + rv
		case MINUS:
			res = lv - rv
		}
	}

	return res
}

func Cmp(leftVal interface{}, rightVal interface{}) int {
	if leftVal == nil && rightVal != nil {
		return -1
	} else if leftVal != nil && rightVal == nil {
		return 1
	} else if leftVal == nil && rightVal == nil {
		return 0
	}

	lT, rT := reflect.TypeOf(leftVal).Kind(), reflect.TypeOf(rightVal).Kind()
	if lT == reflect.String {
		lv, rv := leftVal.(string), rightVal.(string)
		if lv > rv {
			return 1
		} else if lv < rv {
			return -1
		}
		return 0

	} else if lT == reflect.Float64 || rT == reflect.Float64 {
		lv, rv := Val2Float(leftVal), Val2Float(rightVal)
		if lv > rv {
			return 1
		} else if lv < rv {
			return -1
		}
		return 0

	} else {
		lv, rv := Val2Int(leftVal), Val2Int(rightVal)
		if lv > rv {
			return 1
		} else if lv < rv {
			return -1
		}
		return 0
	}
}

func GetFuncType(name string) FuncType {
	if name == "MAX" || name == "MIN" ||
		name == "SUM" || name == "AVG" {
		return AGGREGATE
	} else {
		return NORMAL
	}
}
