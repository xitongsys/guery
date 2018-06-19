package HiveConnector

import (
	"fmt"

	"github.com/xitongsys/guery/FileSystem"
	"github.com/xitongsys/guery/FileSystem/Partition"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/Util"
)

func (self *HiveConnector) Init() (err error) {
	if err := self.setMetadata(); err != nil {
		return err
	}

	if err := self.setTableInfo(); err != nil {
		return err
	}
	return nil
}

func (self *HiveConnector) setMetadata() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(MD_SQL, self.Schema, self.Table, self.Schema, self.Table)
	rows, err := self.db.Query(sqlStr)
	if err != nil {
		return err
	}
	var colName, colType string
	names, types := []string{}, []Type.Type{}

	for rows.Next() {
		rows.Scan(&colName, &colType)
		names = append(names, colName)
		types = append(types, HiveTypeToGueryType(colType))
	}

	if len(names) <= 0 || len(names) != len(types) {
		return fmt.Errorf("[HiveConnector.setMetadata] can't get columns info")
	}

	self.Metadata = Metadata.NewMetadata()
	for i, name := range names {
		t := types[i]
		column := Metadata.NewColumnMetadata(t, "hive", self.Schema, self.Table, name)
		self.Metadata.AppendColumn(column)
	}
	self.Metadata.Reset()
	return nil
}

func (self *HiveConnector) setTableInfo() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(TABLE_INFO_SQL, self.Schema, self.Table)
	rows, err := self.db.Query(sqlStr)
	if err != nil {
		return err
	}
	loc, ft := "", ""
	for rows.Next() {
		rows.Scan(&loc, &ft)
	}
	if loc == "" || ft == "" {
		return fmt.Errorf("HiveConnector: setTableInfo error")
	}
	self.TableLocation = loc
	self.FileType = HiveFileTypeToFileType(ft)
	return nil
}

func (self *HiveConnector) setPartitionInfo() (err error) {
	if err = self.getConn(); err != nil {
		return err
	}
	sqlStr := fmt.Sprintf(PARTITION_MD_SQL, self.Schema, self.Table)
	rows, err := self.db.Query(sqlStr)
	if err != nil {
		return err
	}
	var colName, colType string
	names, types := []string{}, []Type.Type{}
	for rows.Next() {
		rows.Scan(&colName, &colType)
		names = append(names, colName)
		types = append(types, HiveTypeToGueryType(colType))
	}

	md := Metadata.NewMetadata()
	for i, name := range names {
		t := types[i]
		column := Metadata.NewColumnMetadata(t, "hive", self.Schema, self.Table, name)
		md.AppendColumn(column)
	}
	md.Reset()
	self.PartitionInfo = Partition.NewPartitionInfo(md)

	//no partition
	if len(names) <= 0 {
		self.PartitionInfo.FileList, err = FileSystem.List(self.TableLocation)
		for _, f := range self.PartitionInfo.FileList {
			f.FileType = self.FileType
		}
		return err
	}

	//partitioned
	sqlStr = fmt.Sprintf(PARTITION_DATA_SQL, self.Schema, self.Table)
	rows, err = self.db.Query(sqlStr)
	if err != nil {
		return err
	}

	pnum := md.GetColumnNumber()
	partitions := make([]string, pnum)

	location, fileType := "", ""
	i := 0
	for rows.Next() {
		rows.Scan(&location, &fileType, &partitions[i])
		//log.Println("[hiveconninit]====", location, fileType, partitions[i])
		if i == pnum-1 {
			row := Row.NewRow()
			for j := 0; j < pnum; j++ {
				t, err := md.GetTypeByIndex(j)
				if err != nil {
					return err
				}
				row.AppendVals(Type.ToType(partitions[j], t))
				//log.Println("====", partitions[j], t, reflect.TypeOf(Util.ToType(partitions[j], t)))
			}
			self.PartitionInfo.Write(row)

			self.PartitionInfo.Locations = append(self.PartitionInfo.Locations, location)
			ft := HiveFileTypeToFileType(fileType)
			self.PartitionInfo.FileTypes = append(self.PartitionInfo.FileTypes, ft)
			fileList, err := FileSystem.List(location)
			if err != nil {
				return err
			}
			for _, f := range fileList {
				f.FileType = ft
			}
			self.PartitionInfo.FileLists = append(self.PartitionInfo.FileLists, fileList)

			i = 0
		} else {
			i++
		}
	}
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

	db, err := Util.OpenDBConn("mysql", self.Config.GetURI())
	if err != nil {
		self.db = nil
		return err
	}
	self.db = db
	return nil
}
