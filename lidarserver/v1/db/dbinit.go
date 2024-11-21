package lidardb

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var DB *sql.DB
var ctx context.Context
var Qry *Queries

func init() {
	var err error
	if DB == nil {
		DB, err = sql.Open("sqlite3", "1.db")
		if err != nil {
			panic(err)
		}
	}

	// create tables
	ctx = context.Background()
	if _, err := DB.ExecContext(ctx, ddl); err != nil {
		log.Println(err)
	}
	Qry = New(DB)

	Qry.CreateExperiment(ctx,
		CreateExperimentParams{
			Title:     "123123",
			Comments:  "2312312312312",
			Starttime: time.Now(),
			Wavelen:   355.0,
			Vertres:   1500.0,
			Accum:     100,
		})
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
