package lidardb

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeConnection(dbname string) (*Queries, error) {
	dbb, err := sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}

	queries := New(dbb)
	err = queries.CreateTableExperiment(context.Background())
	if err != nil {
		log.Println(err)
	}

	err = queries.CreateTableMeasurement(context.Background())
	if err != nil {
		log.Println(err)
	}
	return queries, nil
}

func CloseConnection(qry *Queries) error {
	if qry != nil {
		qry.db.(*sql.DB).Close()
	}
	return nil
}

var Qry *Queries
