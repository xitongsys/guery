package hive

import (
	"strings"

	"github.com/xitongsys/guery/filesystem"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/row"
)

func HiveTypeToGueryType(ht string) gtype.Type {
	switch strings.ToUpper(ht) {
	case "STRING":
		return gtype.STRING
	case "TINYINT":
		return gtype.INT32
	case "SMALLINT":
		return gtype.INT32
	case "INT":
		return gtype.INT32
	case "BIGINT":
		return gtype.INT64
	case "FLOAT":
		return gtype.FLOAT32
	case "DOUBLE":
		return gtype.FLOAT64
	case "TIMESTAMP":
		return gtype.TIMESTAMP
	case "DATE":
		return gtype.DATE
	default:
		return gtype.UNKNOWNTYPE
	}
}

func HiveFileTypeToFileType(fileType string) filesystem.FileType {
	ss := strings.Split(strings.ToUpper(fileType), ".")
	fileType = ss[len(ss)-1]
	switch fileType {
	case "TEXTINPUTFORMAT":
		return filesystem.CSV
	case "ORCINPUTFORMAT":
		return filesystem.ORC
	case "MAPREDPARQUETINPUTFORMAT":
		return filesystem.PARQUET
	}
	return filesystem.UNKNOWNFILETYPE
}

func HiveTypeConvert(rg *row.RowsGroup) (*row.RowsGroup, error) {
	for i := 0; i < rg.GetColumnsNumber(); i++ {
		t, err := rg.Metadata.GetTypeByIndex(i)
		if err != nil {
			return nil, err
		}
		for j, val := range rg.Vals[i] {
			switch t {
			case gtype.TIMESTAMP:
				switch val.(type) {
				case string:
					s := val.(string)
					if len(s) == 12 { //INT96
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
						rg.Vals[i][j] = gtype.Timestamp{Sec: sec}

					} else {
						rg.Vals[i][j] = gtype.ToTimestamp(val)
					}
				default:
					rg.Vals[i][j] = gtype.ToTimestamp(val)
				}
			case gtype.DATE:
				rg.Vals[i][j] = gtype.ToDate(val)
			}
		}
	}
	return rg, nil
}
