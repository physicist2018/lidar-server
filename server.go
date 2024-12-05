package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"lidarserver.sqlc/app/lidarserver/v2/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		golog.Fatal(err.Error())
	}

	router := http.NewServeMux()
	routes.NewRoutes(router)

	//http.Handle("/", router)
	golog.Info("Сервер запущен на порту 5555")

	// Add logging capability to the router.
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	// Making the server run forever on 5555 port
	// with router.
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         "0.0.0.0:5555",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
