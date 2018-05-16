package Util

import (
	"fmt"
	"time"
)

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
