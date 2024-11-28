package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	lidarserver "lidarserver.sqlc/app/lidarserver/v2/db"
)

func main() {
	golog.SetLevel("info")
	golog.Info("Hello")
	db, err := lidarserver.NewConnection("37.46.134.113:28015", "lidar_admin", "admin")
	if err != nil {
		golog.Error(err)
	}
	if msg, err := db.Ping(); err == nil {
		golog.Info(msg)
	}

	measurements := []*lidarserver.Measurement{
		{
			Id:        "3",
			StartTime: time.Now(),
			StopTime:  time.Now().Add(10 * time.Second),
		},
		{
			Id:        "4",
			StartTime: time.Now(),
			StopTime:  time.Now().Add(10 * time.Second),
		},
	}

	// err = db.CreateMeasurementList(measurements)
	// if err != nil {
	// 	golog.Error(err)
	// }

	err = db.CreateExperiment(&lidarserver.Experiment{
		Data: measurements,
	})
	if err != nil {
		golog.Error(err)
	}

	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:7777",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	golog.Fatal(
		srv.ListenAndServe(),
	)
	db.Close()

}
