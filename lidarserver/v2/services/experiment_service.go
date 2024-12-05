package services

import (
	"context"

	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/bson"
	"lidarserver.sqlc/app/lidarserver/v2/models"
	"lidarserver.sqlc/app/lidarserver/v2/mongodb"
)

type ExperimentService struct{}

var ES ExperimentService

func (e *ExperimentService) FindAll() []models.ExperimentModel {
	var res []models.ExperimentModel
	golog.Info("Entering FindAll")
	clt := mongodb.GetDefaultMongoClientPanic()
	experiment_collection := clt.GetCollection("Experiment")
	cur, err := experiment_collection.Find(context.Background(), bson.D{})
	if err != nil {
		golog.Fatal(err.Error())
	}

	if err = cur.All(context.Background(), &res); err != nil {
		golog.Fatal(err.Error())
	}

	return res
}
