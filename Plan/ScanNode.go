package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type ScanNode struct {
	src            DataSource.TableSource
	ExpressionTree antlr.ParserRuleContext
}

func (self *ScanNode) Filter(val []interface{}) bool {
	switch self.ExpressionTree.(type) {
	case *parser.LogicalNotContext:

	case *parser.LogicalBinaryContext:

	}
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
