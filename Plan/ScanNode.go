package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
)

type ScanNode struct {
	src            DataSource.TableSource
	ExpressionTree antlr.ParserRuleContext
}

func (self *filter) Filter(val []interface{}) bool {

	return true
}

func (self *ScanNode) Children() []PlanNode {
	return nil
}

func (self *ScanNode) Result() DataSource.TableSource {
	ln := len(src.Vals)
	res := DataSource.TableSource{
		Name:            src.Name,
		ColumnNames:     src.ColumnNames,
		ColumnTypes:     src.ColumnTypes,
		ColumnNameIndex: src.ColumnNameIndex,
		Vals:            make([][]interface{}, 0),
	}
	for i := 0; i < ln; i++ {
		if self.Filter(src.Vals[i]) {
			res.Vals = append(res.Vals, src.Vals[i])
		}
	}

	return res

}
