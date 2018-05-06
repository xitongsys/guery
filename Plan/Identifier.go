package Plan

import (
	"fmt"
	"io"

	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type IdentifierNode struct {
	Str         *string
	Digit       *int
	NonReserved *string
}

func NewIdentifierNode(t parser.IIdentifierContext) *IdentifierNode {
	tt := t.(*parser.IdentifierContext)
	res := &IdentifierNode{}
	var (
		str string
		dig int
	)

	if id := tt.IDENTIFIER(); id != nil {
		str = id.GetText()
		res.Str = &str

	} else if qid := tt.QUOTED_IDENTIFIER(); qid != nil {
		str = qid.GetText()
		ln := len(str)
		str = str[1 : ln-1]
		res.Str = &str

	} else if nr := tt.NonReserved(); nr != nil {
		str = nr.GetText()
		res.NonReserved = &str

	} else if did := tt.DIGIT_IDENTIFIER(); did != nil {
		str = did.GetText()
		fmt.Sscanf(str, "%d", &dig)
		res.Digit = &dig
	}
	return res
}

func (self *IdentifierNode) Result(input *Util.RowsBuffer) (interface{}, error) {
	row, err := input.Read()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if self.Digit != nil {
		if *self.Digit >= len(row.Vals) {
			return nil, fmt.Errorf("index out of range")
		}
		return row.Vals[*self.Digit], nil

	} else if self.Str != nil {
		index := input.GetIndex(*self.Str)
		self.Digit = &index
		if *self.Digit >= len(row.Vals) {
			return nil, fmt.Errorf("index out of range")
		}
		return row.Vals[index], nil
	}
	return nil, fmt.Errorf("wrong IdentifierNode")
}

func (self *IdentifierNode) GetText() string {
	if self.Str != nil {
		return *self.Str
	} else if self.Digit != nil {
		return fmt.Sprintf("%d", *self.Digit)
	}
	return ""
}
