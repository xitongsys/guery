package Plan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type NumberNode struct {
	tree      *antlr.Tree
	doubleVal *float64
	intVal    *int
}
