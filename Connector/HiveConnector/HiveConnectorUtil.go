package HiveConnector

import (
	"strings"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
)

func HiveTypeToGueryType(ht string) Type.Type {
	switch strings.ToUpper(ht) {
	case "STRING":
		return Type.STRING
	case "TINYINT":
		return Type.INT32
	case "SMALLINT":
		return Type.INT32
	case "INT":
		return Type.INT32
	case "BIGINT":
		return Type.INT64
	case "FLOAT":
		return Type.FLOAT32
	case "DOUBLE":
		return Type.FLOAT64
	case "TIMESTAMP":
		return Type.TIMESTAMP
	case "DATE":
		return Type.DATE
	default:
		return Type.UNKNOWNTYPE
	}
}

func HiveFileTypeToFileType(fileType string) FileSystem.FileType {
	ss := strings.Split(strings.ToUpper(fileType), ".")
	fileType = ss[len(ss)-1]
	switch fileType {
	case "TEXTINPUTFORMAT":
		return FileSystem.CSV
	case "ORCINPUTFORMAT":
		return FileSystem.ORC
	case "MAPREDPARQUETINPUTFORMAT":
		return FileSystem.PARQUET
	}
	return FileSystem.UNKNOWNFILETYPE
}

func HiveTypeConvert(row *Row.Row, md *Metadata.Metadata) (*Row.Row, error) {
	res := Row.NewRow()
	for i, val := range row.Vals {
		t, err := md.GetTypeByIndex(i)
		if err != nil {
			return res, err
		}
		switch t {
		case Type.TIMESTAMP:
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
					res.AppendVals(Type.ToTimeStamp(sec))

				} else {
					res.AppendVals(Type.ToTimeStamp(val))
				}
			default:
				res.AppendVals(Type.ToTimeStamp(val))
			}
		case Type.DATE:
			res.AppendVals(Type.ToDate(val))
		default:
			res.AppendVals(val)
		}
	}
	return res, nil
}
