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
		t := md.GetTypeByIndex(i)
		switch t {
		case Util.TIMESTAMP:
			switch val.(type) {
			case string:
				s := val.(string)
				if len(s) == 12 { //INT96

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
}
