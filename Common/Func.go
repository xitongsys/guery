package Common

import (
	"fmt"
	"reflect"
)

func Val2Float(val interface{}) float64 {
	v := reflect.ValueOf(val)
	var res float64
	switch v.Kind() {
	case reflect.Int:
		res = float64(v.Int())
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		res = v.Float()
	case reflect.String:
		fmt.Sscanf(v.String(), "%f", &res)
	}
	return res
}

func Val2Int(val interface{}) int64 {
	v := reflect.ValueOf(val)
	var res int64
	switch v.Kind() {
	case reflect.Int:
		res = v.Int()
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		res = int64(v.Float())
	case reflect.String:
		fmt.Sscanf(v.String(), "%d", &res)
	}
	return res
}

func Arithmetic(leftVal interface{}, rightVal interface{}, op Operator) interface{} {
	lT, rT := reflect.TypeOf(leftVal).Kind(), reflect.TypeOf(rightVal).Kind()
	if lT == reflect.Float32 || rT == reflect.Float32 ||
		lT == reflect.Float64 || rT == reflect.Float64 {
		lv, rv := Val2Float(leftVal), Val2Float(rightVal)
		switch op {
		case ASTERISK:
			return lv + rv
		case SLASH:
			return lv / rv
		case PERCENT:
			return nil
		case PLUS:
			return lv + rv
		case MINUS:
			return lv - rv
		}

	} else {
		lv, rv := Val2Int(leftVal), Val2Int(rightVal)
		switch op {
		case ASTERISK:
			return lv + rv
		case SLASH:
			return lv / rv
		case PERCENT:
			return lv % rv
		case PLUS:
			return lv + rv
		case MINUS:
			return lv - rv
		}
	}
	return nil
}
