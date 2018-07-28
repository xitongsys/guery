package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBConn(driverName string, dbURI string) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)
	if db, err = sql.Open(driverName, dbURI); err == nil {
		if err = db.Ping(); err == nil {
			return db, nil
		} else {
			db.Close()
		}
	}
	return db, err
}
