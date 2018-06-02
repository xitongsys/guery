package Plan

import (
	"time"

	"github.com/xitongsys/guery/Util"
)

func NewNowFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "NOW",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.TIMESTAMP, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			return time.Now(), nil
		},
	}
	return res
}
