package services

import (
	"github.com/kataras/golog"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"lidarserver.sqlc/app/lidarserver/v2/dbase"
)

type TestService struct{}

var TS TestService

func (ts *TestService) GetTables() []string {
	client := dbase.GetConnection()

	curs, err := r.DB(dbase.DB_NAME).TableList().Run(client.Session)
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
