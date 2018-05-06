package Util

type Row struct {
	Key  string
	Vals []interface{}
}

func NewRow(vals ...interface{}) *Row {
	res := &Row{}
	return res
}

func (self *Row) AppendVals(vals ...interface{}) *Row {
	self.Vals = append(self.Vals, vals...)
	return self
}
