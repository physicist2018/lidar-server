package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"lidarserver.sqlc/app/lidarserver/v2/routes"
)

// Logging middleware for logging each IO operation
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		golog.Info(fmt.Sprintf("%s\t%s\t%.3f ms", r.Method, r.RequestURI, float64(time.Since(start).Microseconds())/1000.0))
	})
}

func main() {

	port := flag.Int("port", 5432, "Порт, на котором будет вращаться сервер.")
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		golog.Fatal(err.Error())
	}

	router := http.NewServeMux()
	routes.NewRoutes(router)

	golog.Info(fmt.Sprintf("Server started at port: %d", *port))

	// Add logging capability to the router.
	loggedRouter := Logging(router)

	// Making the server run forever on 5555 port
	// with router.
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf("0.0.0.0:%d", *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
