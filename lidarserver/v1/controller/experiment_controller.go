package controller

import (
	"context"

	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
)

func GetExperimentById(idd int) lidardb.Experiment {
	exp, err := lidardb.Qry.GetExperimentById(context.Background(), int64(idd))
	if err != nil {
		panic(err)
	}
	return exp
}

func GetAllExperients() []lidardb.Experiment {
	exps, err := lidardb.Qry.GetAllExperiments(context.Background())
	if err != nil {
		panic(err)
	}

	return exps
}
