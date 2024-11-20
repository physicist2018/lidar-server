package lidardb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeConnection(dbname string) (*Queries, error) {
	dbb, err := sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}

	queries := New(dbb)
	return queries, nil
}

func CloseConnection(qry *Queries) error {
	if qry != nil {
		qry.db.(*sql.DB).Close()
	}
	return nil
}

var Qry *Queries

func init() {
	Qry, _ = InitializeConnection("1.db")
}
