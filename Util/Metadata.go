package Util

import (
	"fmt"
	"strings"
)

type Metadata struct {
	Columns   []*ColumnMetadata
	Keys      []*ColumnMetadata
	ColumnMap map[string]int
}

func (self *Metadata) Reset() {
	self.ColumnMap = map[string]int{}
	for i, col := range self.Columns {
		name := col.ColumnName
		self.ColumnMap[name] = i

		name = col.Table + "." + name
		self.ColumnMap[name] = i

		name = col.Schema + "." + name
		self.ColumnMap[name] = i

		name = col.Catalog + "." + name
		self.ColumnMap[name] = i
	}
}

func (self *Metadata) GetColumnNames() []string {
	res := []string{}
	for _, c := range self.Columns {
		res = append(res, c.GetName())
	}
	return res
}

func (self *Metadata) GetColumnTypes() []Type {
	res := []Type{}
	for _, c := range self.Columns {
		res = append(res, c.ColumnType)
	}
	return res
}

func (self *Metadata) Copy() *Metadata {
	res := NewMetadata()
	for _, c := range self.Columns {
		res.Columns = append(res.Columns, c.Copy())
	}
	for _, k := range self.Keys {
		res.Keys = append(res.Keys, k.Copy())
	}
	res.Reset()
	return res
}

func (self *Metadata) Rename(name string) {
	for _, c := range self.Columns {
		c.Table = name
	}
	self.Reset()
}

func (self *Metadata) GetColumnNumber() int {
	return len(self.Columns)
}

func (self *Metadata) GetTypeByIndex(index int) (Type, error) {
	if index >= len(self.Columns) {
		return UNKNOWNTYPE, fmt.Errorf("index out of range")
	}
	return self.Columns[index].ColumnType, nil
}

func (self *Metadata) GetTypeByName(name string) (Type, error) {
	index, ok := self.ColumnMap[name]
	if !ok {
		return UNKNOWNTYPE, fmt.Errorf("unknown column name: %v", name)
	}
	return self.GetTypeByIndex(index)
}

func (self *Metadata) GetIndexByName(name string) (int, error) {
	index, ok := self.ColumnMap[name]
	if !ok {
		return -1, fmt.Errorf("unknown column name: %v", name)
	}
	return index, nil
}

func (self *Metadata) AppendColumn(column *ColumnMetadata) {
	self.Columns = append(self.Columns, column)
	self.Reset()
}

func (self *Metadata) AppendKey(key *ColumnMetadata) {
	self.Keys = append(self.Keys, key)
}

func (self *Metadata) AppendKeyByType(t Type) {
	k := &ColumnMetadata{
		ColumnType: t,
	}
	self.Keys = append(self.Keys, k)
}

func (self *Metadata) ClearKeys() {
	self.Keys = []*ColumnMetadata{}
}

func (self *Metadata) DeleteColumnByIndex(index int) {
	ln := len(self.Columns)
	if index < 0 || index >= ln {
		return
	}
	self.Columns = append(self.Columns[:index], self.Columns[index+1:]...)
}

func (self *Metadata) SelectColumns(columns []string) *Metadata {
	res := NewMetadata()
	rec := map[int]bool{}
	for _, c := range columns {
		index, err := self.GetIndexByName(c)
		if err != nil {
			continue
		}
		if _, ok := rec[index]; !ok {
			rec[index] = true
			res.Columns = append(res.Columns, self.Columns[index].Copy())
		}
	}
	res.Reset()
	return res
}

func (self *Metadata) Contains(columns []string) bool {
	for _, c := range columns {
		if _, err := self.GetIndexByName(c); err != nil {
			return false
		}
	}
	return true
}

func NewMetadata() *Metadata {
	return &Metadata{
		Columns:   []*ColumnMetadata{},
		Keys:      []*ColumnMetadata{},
		ColumnMap: map[string]int{},
	}
}

func SplitTableName(name string) (catalog, schema, table string) {
	catalog, schema, table = "default", "default", "default"
	names := strings.Split(name, ".")
	ln := len(names)
	if ln >= 1 {
		table = names[ln-1]
	}
	if ln >= 2 {
		schema = names[ln-2]
	}
	if ln >= 3 {
		catalog = names[ln-3]
	}
	return
}

func JoinMetadata(mdl, mdr *Metadata) *Metadata {
	res := NewMetadata()
	for _, c := range mdl.Columns {
		res.Columns = append(res.Columns, c.Copy())
	}
	for _, k := range mdl.Keys {
		res.Keys = append(res.Keys, k.Copy())
	}
	for _, c := range mdr.Columns {
		res.Columns = append(res.Columns, c.Copy())
	}
	for _, k := range mdr.Keys {
		res.Keys = append(res.Keys, k.Copy())
	}
	res.Reset()
	return res
}

/*Json metadata struct
{
    "Catalog":"FILE",
    "Schema": "INFO",
    "Table": "STUDENT",
    "ColumnNames": ["ID","NAME","AGE"],
    "ColumnTypes": ["INT64","STRING","INT32"],
    "KeyTypes": ["INT64"]
}
*/

type JsonMetadata struct {
	Catalog     string
	Schema      string
	Table       string
	ColumnNames []string
	ColumnTypes []string
	KeyTypes    []string
}

func NewMetadataFromJsonMetadata(jm *JsonMetadata) (*Metadata, error) {
	res := NewMetadata()
	if len(jm.ColumnNames) != len(jm.ColumnTypes) {
		return res, fmt.Errorf("JsonMetadata format error")
	}

	for i := 0; i < len(jm.ColumnNames); i++ {
		col := &ColumnMetadata{
			Catalog:    jm.Catalog,
			Schema:     jm.Schema,
			Table:      jm.Table,
			ColumnName: jm.ColumnNames[i],
			ColumnType: TypeNameToType(jm.ColumnTypes[i]),
		}
		res.AppendColumn(col)
	}

	for i := 0; i < len(jm.KeyTypes); i++ {
		key := &ColumnMetadata{
			ColumnType: TypeNameToType(jm.KeyTypes[i]),
		}
		res.AppendKey(key)
	}

	res.Reset()
	return res, nil
}
