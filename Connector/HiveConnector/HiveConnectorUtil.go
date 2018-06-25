package HiveConnector

import (
	"strings"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Split"
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

func HiveTypeConvert(sp *Split.Split, md *Metadata.Metadata, indexes []int) (*Split.Split, error) {
	for j := 0; j < sp.GetRowsNumber(); j++ {
		for i := 0; i < sp.GetColumnNumber(); i++ {
			t, err := md.GetTypeByIndex(indexes[i])
			if err != nil {
				return sp, err
			}
			val := sp.Values[i][j]
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
						sp.Values[i][j] = Type.ToTimeStamp(sec)

					} else {
						sp.Values[i][j] = Type.ToTimeStamp(val)
					}
				default:
					sp.Values[i][j] = Type.ToTimeStamp(val)
				}
			case Type.DATE:
				sp.Values[i][j] = Type.ToDate(val)

			}
		}
	}
	return sp, nil
}
