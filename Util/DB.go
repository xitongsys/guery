package Util

func OpenDBConn(driverName string, dbConnInfo *DBConnInfoType) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)
	retry := 0
	for retry < Conf.QueueConf.MaxRetry {
		if db, err = sql.Open(driverName, Conf.GetDatabaseConnStr(dbConnInfo)); err != nil {
			time.Sleep(5 * time.Second)
			retry++
			continue
		}
		err = db.Ping()
		if err != nil {
			time.Sleep(5 * time.Second)
			db.Close()
			retry++
		} else {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
