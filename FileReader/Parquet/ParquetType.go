package Parquet

import (
	. "github.com/xitongsys/parquet-go/ParquetType"
	"github.com/xitongsys/parquet-go/parquet"
)

func ParquetTypeToGueryType(src interface{}, pT *parquet.Type, cT *parquet.ConvertedType) interface{} {
	if src == nil {
		return nil
	}
	if cT == nil {
		if *pT == parquet.Type_BOOLEAN {
			return bool(src.(BOOLEAN))
		} else if *pT == parquet.Type_INT32 {
			return int32(src.(INT32))
		} else if *pT == parquet.Type_INT64 {
			return int64(src.(INT64))
		} else if *pT == parquet.Type_INT96 {
			return string(src.(INT96))
		} else if *pT == parquet.Type_FLOAT {
			return float32(src.(FLOAT))
		} else if *pT == parquet.Type_DOUBLE {
			return float64(src.(DOUBLE))
		} else if *pT == parquet.Type_BYTE_ARRAY {
			return string(src.(BYTE_ARRAY))
		} else if *pT == parquet.Type_FIXED_LEN_BYTE_ARRAY {
			return string(src.(FIXED_LEN_BYTE_ARRAY))
		}
		return nil
	}

	if *cT == parquet.ConvertedType_UTF8 {
		return string(src.(BYTE_ARRAY))
	} else if *cT == parquet.ConvertedType_INT_8 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_INT_16 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_INT_32 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_INT_64 {
		return int64(src.(INT64))
	} else if *cT == parquet.ConvertedType_UINT_8 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_UINT_16 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_UINT_32 {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_UINT_64 {
		return int64(src.(INT64))
	} else if *cT == parquet.ConvertedType_DATE {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_TIME_MILLIS {
		return int32(src.(INT32))
	} else if *cT == parquet.ConvertedType_TIME_MICROS {
		return int64(src.(INT64))
	} else if *cT == parquet.ConvertedType_TIMESTAMP_MILLIS {
		return int64(src.(INT64))
	} else if *cT == parquet.ConvertedType_TIMESTAMP_MICROS {
		return int64(src.(INT64))
	} else if *cT == parquet.ConvertedType_INTERVAL {
		return string(src.(FIXED_LEN_BYTE_ARRAY))
	} else if *cT == parquet.ConvertedType_DECIMAL {
		if *pT == parquet.Type_INT32 {
			return int32(src.(INT32))
		} else if *pT == parquet.Type_INT64 {
			return int64(src.(INT64))
		} else if *pT == parquet.Type_FIXED_LEN_BYTE_ARRAY {
			return string(src.(FIXED_LEN_BYTE_ARRAY))
		} else {
			return string(src.(BYTE_ARRAY))
		}
	} else {
		return nil
	}
}
