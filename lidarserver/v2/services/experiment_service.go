package services

import (
	"context"

	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lidarserver.sqlc/app/lidarserver/v2/models"
	"lidarserver.sqlc/app/lidarserver/v2/mongodb"
)

type ExperimentService struct{}

var ES ExperimentService

func (e *ExperimentService) FindAll() []models.ExperimentModel {
	var res []models.ExperimentModel
	golog.Info("-- Entering FindAll")
	clt := mongodb.GetDefaultMongoClientPanic()
	experiment_collection := clt.GetCollection("Experiment")
	cur, err := experiment_collection.Find(context.Background(), bson.D{})
	if err != nil {
		golog.Error(err.Error())
		return nil
	}

	if err = cur.All(context.Background(), &res); err != nil {
		golog.Fatal(err.Error())
	}
	golog.Info("-- Leaving FindAll")
	return res
}

func (e *ExperimentService) FindOne(id string) *models.ExperimentModel {

	golog.Info("-- Entering FindOne with id: ", id)

	clt := mongodb.GetDefaultMongoClientPanic()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	experiment_collection := clt.GetCollection("Experiment")
	result := experiment_collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result == nil {
		golog.Fatal("Error getting document")
		return nil
	}
	res := &models.ExperimentModel{}
	result.Decode(res)
	golog.Info("-- Leaving FindOne")

	return res
}

func (e *ExperimentService) DeleteOne(id string) error {
	golog.Info("-- Entering DeleteOne with id: ", id)

	clt := mongodb.GetDefaultMongoClientPanic()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		golog.Error(err.Error())
		return err
	}
	experiment_collection := clt.GetCollection("Experiment")

	result, err := experiment_collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		golog.Error(err.Error())

		return err
	}

	golog.Info(result)
	return nil
}
