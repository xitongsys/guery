package HiveConnector

import (
	"strings"

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

func HiveFileTypeToSimpleType(fileType string) string {
	ss := strings.Split(strings.ToUpper(fileType), ".")
	fileType = ss[len(ss)-1]
	switch fileType {
	case "TEXTINPUTFORMAT":
		return "CSV"
	case "ORCINPUTFORMAT":
		return "ORC"
	case "MAPREDPARQUETINPUTFORMAT":
		return "PARQUET"
	}
	return ""
}
