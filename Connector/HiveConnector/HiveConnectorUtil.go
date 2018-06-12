package HiveConnector

import (
	"strings"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Util"
)

func HiveTypeToGueryType(ht string) Util.Type {
	switch strings.ToUpper(ht) {
	case "STRING":
		return Util.STRING
	case "TINYINT":
		return Util.INT32
	case "SMALLINT":
		return Util.INT32
	case "INT":
		return Util.INT32
	case "BIGINT":
		return Util.INT64
	case "FLOAT":
		return Util.FLOAT32
	case "DOUBLE":
		return Util.FLOAT64
	case "TIMESTAMP":
		return Util.TIMESTAMP
	case "DATE":
		return Util.DATE
	default:
		return Util.UNKNOWNTYPE
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

func HiveTypeConvert(row *Util.Row, md *Util.Metadata) (*Util.Row, error) {
	res := Util.NewRow()
	for i, val := range row.Vals {
		t, err := md.GetTypeByIndex(i)
		if err != nil {
			return res, err
		}
		switch t {
		case Util.TIMESTAMP:
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
					res.AppendVals(Util.ToTimeStamp(sec))

				} else {
					res.AppendVals(Util.ToTimeStamp(val))
				}
			default:
				res.AppendVals(Util.ToTimeStamp(val))
			}
		case Util.DATE:
			res.AppendVals(Util.ToDate(val))
		default:
			res.AppendVals(val)
		}
	}
	return res, nil
}
