package lidardb

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func InitializeConnection(dbname string) (*Queries, error) {
	dbb, err := sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}

	if res, err := dbb.ExecContext(context.Background(), ddl); err != nil {
		log.Panicln(err, res)
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

// func init() {
// 	Qry, _ = InitializeConnection("1.db")
// 	r, err := Qry.CreateExperiment(context.Background(),
// 		CreateExperimentParams{
// 			Title:     "New experiment",
// 			Comments:  "Simple experiment",
// 			StartTime: time.Now(),
// 			Wavelen:   532.0,
// 			VertRes:   1500.0,
// 			Accum:     1,
// 		})
// 	log.Println(r, err)
// }
