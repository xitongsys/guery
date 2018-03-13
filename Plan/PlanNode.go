package Plan

import (
	"github.com/xitongsys/guery/DataSource"
)

type PlanNode interface {
	Children() []PlanNode
	Result() DataSource.TableSource
}
