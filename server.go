package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
	"lidarserver.sqlc/app/lidarserver/v1/routes"
)

//go:embed schema.sql
var ddl string

func main() {
	dbname := "lidar.db"
	if len(os.Args) == 2 {
		dbname = os.Args[1]
	}

	var err error
	lidardb.Qry, err = lidardb.InitializeConnection(dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer lidardb.CloseConnection(lidardb.Qry)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	routes.MakeRoutes(router)

	router.Run("localhost:7777")
}
