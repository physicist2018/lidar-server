package dbase

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"lidarserver.sqlc/app/lidarserver/v2/models"
)

type DBClient struct {
	Session *r.Session
}

var client *DBClient
var TABLE_NAME = ""
var DB_NAME = ""

func NewConnectionPanic() *DBClient {
	session, err := r.Connect(r.ConnectOpts{
		Address:  os.Getenv("RETHINKDB_URI"),
		Password: os.Getenv("RETHINKDB_PASS"),
		Username: os.Getenv("RETHINKDB_USER"),
		Database: os.Getenv("RETHINKDB_DB"),
	})

	if err != nil {
		golog.Error(err.Error())
		panic(err)
	}

	return &DBClient{
		Session: session,
	}
}

func GetConnection() *DBClient {
	if client != nil {
		return client
	}

	client = NewConnectionPanic()
	return client
}

func GetTables() []string {
	client := GetConnection()

	curs, err := r.DB(os.Getenv(DB_NAME)).TableList().Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return []string{}
	}

	var tbls []string
	err = curs.All(&tbls)

	if err != nil {
		golog.Error(err.Error())
		return []string{}
	}
	return tbls
}

func InsertExperiment(e *models.ExperimentModel) error {
	client := GetConnection()

	cur, err := r.DB(DB_NAME).Table(TABLE_NAME).Insert(e).Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return err
	}
	var res []map[string]interface{}
	cur.All(&res)
	golog.Info(res[0])
	return nil
}

func GetAllExperiments() ([]models.ExperimentModel, error) {
	client := GetConnection()
	cur, err := r.DB(DB_NAME).Table(TABLE_NAME).Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return nil, err
	}
	var ret []models.ExperimentModel
	err = cur.All(&ret)
	if err != nil {
		golog.Error(err.Error())
		return nil, err
	}
	return ret, nil
}

func init() {
	godotenv.Load()
	TABLE_NAME = os.Getenv("RETHINKDB_TBLNAME")
	DB_NAME = os.Getenv("RETHINKDB_DB")
	client = GetConnection()
}
