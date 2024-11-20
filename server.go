package main

import (
	_ "embed"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
	"lidarserver.sqlc/app/lidarserver/v1/routes"
)

//go:embed schema.sql
var ddl string

func main() {
	qry, err := lidardb.InitializeConnection("1.db")
	if err != nil {
		log.Fatal(err)
	}
	defer lidardb.CloseConnection(qry)

	router := gin.Default()

	routes.MakeRoutes(router)

	router.Run(":7777")
}
