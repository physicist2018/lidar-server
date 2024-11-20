package controller

import (
	"context"

	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
)

func GetAllExperiments(idd int) lidardb.Experiment {
	exp, err := lidardb.Qry.GetExperimentById(context.Background(), int64(idd))
	if err != nil {
		panic(err)
	}
	return exp
}
