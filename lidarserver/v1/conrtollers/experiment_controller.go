package conrtollers

import (
	"context"

	lidardb "lidarserver.sqlc/app/lidarserver/v1/db"
)

func GetExperimentById(id int64) (lidardb.Experiment, error) {
	return lidardb.Qry.GetExperimentById(context.Background(), id)
}
