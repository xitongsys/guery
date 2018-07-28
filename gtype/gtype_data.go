package gtype

import (
	"fmt"
	"time"
)

//////////////////
type Type int32

const (
	UNKNOWNTYPE Type = iota
	STRING
	FLOAT64
	FLOAT32
	INT64
	INT32
	BOOL
	TIMESTAMP
	DATE
)

func (self Type) String() string {
	switch self {
	case STRING:
		return "STRING"
	case FLOAT32:
		return "FLOAT32"
	case FLOAT64:
		return "FLOAT64"
	case INT64:
		return "INT64"
	case INT32:
		return "INT32"
	case BOOL:
		return "BOOL"
	case TIMESTAMP:
		return "TIMESTAMP"
	case DATE:
		return "DATE"
	}
	return "UNKNOWNTYPE"
}

func TypeNameToType(name string) Type {
	switch name {
	case "STRING":
		return STRING
	case "FLOAT32":
		return FLOAT32
	case "FLOAT64":
		return FLOAT64
	case "INT64":
		return INT64
	case "INT32":
		return INT32
	case "BOOL":
		return BOOL
	case "TIMESTAMP":
		return TIMESTAMP
	case "DATE":
		return DATE
	default:
		return UNKNOWNTYPE
	}
}

/////////////////////////
func TypeOf(v interface{}) Type {
	switch v.(type) {
	case bool:
		return BOOL
	case int32:
		return INT32
	case int64:
		return INT64
	case float32:
		return FLOAT32
	case float64:
		return FLOAT64
	case string:
		return STRING
	case Timestamp:
		return TIMESTAMP
	case Date:
		return DATE
	default:
		return UNKNOWNTYPE
	}
}

//STRING////////////
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//INT//////////////
func ToInt32(v interface{}) int32 {
	var res int32
	switch v.(type) {
	case bool:
		if v.(bool) {
			res = 1
		}
	case int32:
		res = v.(int32)
	case int64:
		res = int32(v.(int64))
	case float32:
		res = int32(v.(float32))
	case float64:
		res = int32(v.(float64))
	case string:
		fmt.Sscanf(v.(string), "%d", &res)
	case Timestamp:
		res = int32(v.(Timestamp).Sec)
	case Date:
		res = int32(v.(Date).Sec)
	}
	return res
}

func ToInt64(v interface{}) int64 {
	var res int64
	switch v.(type) {
	case bool:
		if v.(bool) {
			res = 1
		}
	case int32:
		res = int64(v.(int32))
	case int64:
		res = v.(int64)
	case float32:
		res = int64(v.(float32))
	case float64:
		res = int64(v.(float64))
	case string:
		fmt.Sscanf(v.(string), "%d", &res)
	case time.Time:
		res = int64(v.(Timestamp).Sec)
	case Date:
		res = int64(v.(Date).Sec)
	}
	return res
}

//FLOAT/////////////
func ToFloat32(v interface{}) float32 {
	var res float32
	switch v.(type) {
	case bool:
		if v.(bool) {
			res = 1.0
		}
	case int32:
		res = float32(v.(int32))
	case int64:
		res = float32(v.(int64))
	case float32:
		res = v.(float32)
	case float64:
		res = float32(v.(float64))
	case string:
		fmt.Sscanf(v.(string), "%f", &res)
	case time.Time:
		res = float32(v.(Timestamp).Sec)
	case Date:
		res = float32(v.(Date).Sec)
	}
	return res
}

func ToFloat64(v interface{}) float64 {
	var res float64
	switch v.(type) {
	case bool:
		if v.(bool) {
			res = 1.0
		}
	case int32:
		res = float64(v.(int32))
	case int64:
		res = float64(v.(int64))
	case float32:
		res = float64(v.(float32))
	case float64:
		res = v.(float64)
	case string:
		fmt.Sscanf(v.(string), "%f", &res)
	case time.Time:
		res = float64(v.(Timestamp).Sec)
	case Date:
		res = float64(v.(Date).Sec)
	}
	return res
}

//TIME////////////////
func ToTimestamp(v interface{}) Timestamp {
	var res Timestamp
	var err error
	var sec int64
	switch v.(type) {
	case bool:
	case int32:
		sec = int64(v.(int32))
	case int64:
		sec = int64(v.(int64))
	case float32:
		sec = int64(v.(float32))
	case float64:
		sec = int64(v.(float64))
	case string:
		var t time.Time
		if t, err = time.Parse(time.RFC3339, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.UnixDate, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RubyDate, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC822, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC822Z, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC850, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC1123, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC1123Z, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC3339, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC3339Nano, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse("2006-01-02", v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse("2006-01-02 15:04:05", v.(string)); err == nil {
			sec = t.Unix()
		}

	case Timestamp:
		sec = v.(Timestamp).Sec
	case Date:
		sec = v.(Date).Sec
	}
	res.Sec = sec
	return res
}

//DATE///////////////
func ToDate(v interface{}) Date {
	var res Date
	var err error
	var sec int64
	switch v.(type) {
	case bool:
	case int32:
		sec = int64(v.(int32))
	case int64:
		sec = int64(v.(int64))
	case float32:
		sec = int64(v.(float32))
	case float64:
		sec = int64(v.(float64))
	case string:
		var t time.Time
		if t, err = time.Parse(time.RFC3339, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.UnixDate, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RubyDate, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC822, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC822Z, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC850, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC1123, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC1123Z, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC3339, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse(time.RFC3339Nano, v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse("2006-01-02", v.(string)); err == nil {
			sec = t.Unix()
		} else if t, err = time.Parse("2006-01-02 15:04:05", v.(string)); err == nil {
			sec = t.Unix()
		}

	case Timestamp:
		sec = v.(Timestamp).Sec
	case Date:
		sec = v.(Date).Sec
	}
	res.Sec = sec
	return res
}

//BOOL/////////////
func ToBool(v interface{}) bool {
	var res bool
	switch v.(type) {
	case bool:
		res = v.(bool)
	case int32:
		if v.(int32) != 0 {
			res = true
		}
	case int64:
		if v.(int64) != 0 {
			res = true
		}
	case float32:
		if v.(float32) != 0 {
			res = true
		}
	case float64:
		if v.(float64) != 0 {
			res = true
		}
	case string:
		if v.(string) != "" {
			res = true
		}
	case time.Time:
		res = true
	case Date:
		res = true
	}
	return res
}

func ToType(v interface{}, t Type) interface{} {
	var res interface{}
	switch t {
	case BOOL:
		res = ToBool(v)
	case INT32:
		res = ToInt32(v)
	case INT64:
		res = ToInt64(v)
	case FLOAT32:
		res = ToFloat32(v)
	case FLOAT64:
		res = ToFloat64(v)
	case STRING:
		res = ToString(v)
	case TIMESTAMP:
		res = ToTimestamp(v)
	case DATE:
		res = ToDate(v)
	}
	return res
}

func ToSameType(va interface{}, vb interface{}) (interface{}, interface{}) {
	ta, tb := TypeOf(va), TypeOf(vb)
	var t Type
	if tb >= ta {
		t = ta
	} else {
		t = tb
	}
	return ToType(va, t), ToType(vb, t)
}
