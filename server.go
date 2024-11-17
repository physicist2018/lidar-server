package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
	"lidarserver.sqlc/app/lidarserver/v1/routes"
)

//go:embed schema.sql
var ddl string

func run() error {

	ctx := context.Background()
	db, err := sql.Open("sqlite3", "1.db")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Println(err)
	}

	queries := lidardb.New(db)

	res, err := queries.CreateExperiment(ctx,
		lidardb.CreateExperimentParams{
			Title:     "123123",
			Comments:  "2312312312312",
			Starttime: time.Now(),
			Wavelen:   355.0,
			Vertres:   1500.0,
			Accum:     100,
		})
	if err != nil {
		return err
	}

	log.Println(res)
	res11, err1 := queries.GetExperiment(ctx, 1)
	// list all authors
	if err1 != nil {
		return err
	}
	log.Println(res11)
	res2, err := queries.GetAllExperiments(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range res2 {
		log.Println(v.ID, v.Starttime, v.Title, v.Comments, v.Wavelen)
	}
	return nil
}

func main() {
	mux := http.NewServeMux()

	if err := run(); err != nil {
		log.Fatal(err)
	}
	routes.MakeRoutes(mux)
	http.ListenAndServe(":7777", mux)
}
