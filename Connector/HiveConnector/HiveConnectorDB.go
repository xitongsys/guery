package HiveConnector

import (
	"fmt"

	"github.com/xitongsys/guery/Util"
)

var MDSQL = `
select P.PKEY_NAME as name, P.PKEY_TYPE as type from 
TBLS as T 
join DBS as D on D.DB_ID=T.DB_ID
join PARTITION_KEYS as P on T.TBL_ID=P.TBL_ID
where D.NAME='%s' and  T.TBL_NAME='%s'
order by T.TBL_ID, P.INTEGER_IDX;
`

func (self *HiveConnector) setMetadata() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(MSDSQL, self.Schema, self.Table)
	rows, err := self.db.Query()
	if err != nil {
		return err
	}
	var colName, colType string
	names, types := []string{}, []Util.Type{}
	for rows.Next() {
		rows.Scan(&colName, colType)
		names = append(names, colName)
		types = append(types, HiveTypeToGueryType(colType))
	}
	self.Metadata = Util.NewMetadata()
	for i, name := range names {
		t := types[i]
		column := Util.NewColumnMetadata(t, "HIVE", self.Schema, self.Table, name)
		self.Metadata.AppendColumn(column)
	}
	self.Metadata.Reset()
	return nil
}

func (self *HiveConnector) getConn() error {
	if self.db != nil {
		if err := self.db.Ping(); err != nil {
			self.db.Close()
			self.db = nil
		} else {
			return nil
		}
	}

	db, err := Util.OpenDBConn(self.Config.GetURI())
	if err != nil {
		self.db = nil
		return err
	}
	self.db = db
	return nil
}
