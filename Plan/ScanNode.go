package Plan

import (
	"github.com/xitongsys/guery/DataSource"
)

type ScanNode struct {
	src    DataSource.TableSource
	filter Filter
}

func (self *ScanNode) Children() []PlanNode {
	return nil
}

func (self *ScanNode) Result() DataSource.TableSource {

}
