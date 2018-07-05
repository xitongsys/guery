package Parquet

import (
	"time"

	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/parquet-go/parquet"
)

func ParquetTypeToGueryType(src interface{}, pT *parquet.Type, cT *parquet.ConvertedType, gt Type.Type) interface{} {
	if src == nil {
		return nil
	}
	if gt == Type.TIMESTAMP {
		if *pT == parquet.Type_INT96 {
			s := src.(string)
			//first 8 byte is a int64 value for nanoseconds of the day
			//last 4 byte is a int32 value for julian day
			nanosec := int64(0)
			base := int64(1)
			for i := 0; i < 8; i++ {
				nanosec = nanosec + int64(s[i])*base
				base = base * 256
			}
			base = 1
			day := int64(0) - 2440588 //jd(1970-01-01)=2440588
			for i := 8; i < 12; i++ {
				day = day + int64(s[i])*base
				base = base * 256
			}
			sec := nanosec/1000000000 + day*3600*24
			src = int64(sec)
		}
		src = time.Unix(src.(int64), 0)
		//src = Type.ToType(src, gt)

	} else if gt == Type.DATE {
		src = Type.ToType(src, gt)
	}

	return src
}
