package Plan

import (
	"strings"

	"github.com/xitongsys/guery/parser"
)

type QualifiedNameNode struct {
	Name string
}

func NewQulifiedNameNode(t parser.IQualifiedNameContext) *QualifiedNameNode {
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
