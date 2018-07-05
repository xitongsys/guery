package Orc

import (
	"time"

	"github.com/scritchley/orc"
	"github.com/scritchley/orc/proto"
	"github.com/xitongsys/guery/Type"
)

func OrcTypeToGueryType(src interface{}, oT proto.Type_Kind) interface{} {
	switch oT {
	case proto.Type_BOOLEAN:
		return src
	case proto.Type_BYTE:
		return nil
	case proto.Type_SHORT, proto.Type_INT:
		return int32(src.(int64))
	case proto.Type_LONG:
		return src
	case proto.Type_FLOAT:
		return src
	case proto.Type_DOUBLE:
		return src
	case proto.Type_STRING, proto.Type_VARCHAR, proto.Type_CHAR:
		return src
	case proto.Type_BINARY:
		return string(src.([]byte))
	case proto.Type_TIMESTAMP:
		return Type.Timestamp{Sec: src.(time.Time).Unix()}
	case proto.Type_LIST:
		return nil
	case proto.Type_MAP:
		return nil
	case proto.Type_STRUCT:
		return nil
	case proto.Type_UNION:
		return nil
	case proto.Type_DECIMAL:
		return nil
	case proto.Type_DATE:
		return Type.Date{Sec: src.(orc.Date).Unix()}
	}
	return nil
}
