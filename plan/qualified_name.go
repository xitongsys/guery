package plan

import (
	"strings"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
)

type QualifiedNameNode struct {
	Name string
}

func NewQulifiedNameNode(runtime *Config.ConfigRuntime, t parser.IQualifiedNameContext) *QualifiedNameNode {
	res := &QualifiedNameNode{}
	tt := t.(*parser.QualifiedNameContext)
	ids := tt.AllIdentifier()
	names := []string{}
	for i := 0; i < len(ids); i++ {
		id := ids[i].(*parser.IdentifierContext)
		names = append(names, id.GetText())
	}
	res.Name = strings.Join(names, ".")
	return res
}

func (self *QualifiedNameNode) Result() string {
	return self.Name
}

func (self *QualifiedNameNode) Init(md *Metadata.Metadata) error {
	return nil
}
