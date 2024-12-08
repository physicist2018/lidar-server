package services

import (
	"github.com/google/uuid"
	"github.com/kataras/golog"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"lidarserver.sqlc/app/lidarserver/v2/dbase"
	"lidarserver.sqlc/app/lidarserver/v2/models"
)

type ExperimentService struct{}

var ES ExperimentService

func (e *ExperimentService) FindAll() []models.ExperimentModel {

	client := dbase.GetConnection()
	cur, err := r.DB(dbase.DB_NAME).Table(dbase.TABLE_NAME).Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	var ret []models.ExperimentModel
	err = cur.All(&ret)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	return ret
}

func (e *ExperimentService) FindOne(id string) *models.ExperimentModel {
	client := dbase.GetConnection()
	cur, err := r.DB(dbase.DB_NAME).Table(dbase.TABLE_NAME).Get(id).Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	ret := &models.ExperimentModel{}
	err = cur.One(ret)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}

	return ret
}

func (e *ExperimentService) UpdateOne(id string, m *models.ExperimentModel) *models.ResponseModel {
	client := dbase.GetConnection()
	cur, err := r.DB(dbase.DB_NAME).Table(dbase.TABLE_NAME).Get(id).Update(m).Run(client.Session)

	if err != nil {
		golog.Error(err.Error())
		return nil
	}

	res := &models.ResponseModel{}
	err = cur.One(res)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	return res
}

func (e *ExperimentService) DeleteOne(id string) *models.ResponseModel {
	client := dbase.GetConnection()
	cur, err := r.DB(dbase.DB_NAME).Table(dbase.TABLE_NAME).Get(id).Delete().Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	res := &models.ResponseModel{}
	err = cur.One(&res)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	return res
}

func (e *ExperimentService) CreateOne(m *models.ExperimentModel) *models.ResponseModel {

	client := dbase.GetConnection()

	//prepare profiles, initialise id
	for i := range m.Dat {
		m.Dat[i].ID = uuid.NewString()
	}

	for i := range m.Dak {
		m.Dak[i].ID = uuid.NewString()
	}
	//
	cur, err := r.DB(dbase.DB_NAME).Table(dbase.TABLE_NAME).Insert(m).Run(client.Session)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	res := &models.ResponseModel{}
	err = cur.One(&res)
	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	golog.Info(res)
	return res
}

// func (e *ExperimentService) FindAll() []models.ExperimentModel {
// 	var res []models.ExperimentModel
// 	golog.Info("-- Entering FindAll")
// 	clt := mongodb.GetDefaultMongoClientPanic()
// 	experiment_collection := clt.GetCollection("Experiment")
// 	cur, err := experiment_collection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		golog.Error(err.Error())
// 		return nil
// 	}

// 	if err = cur.All(context.Background(), &res); err != nil {
// 		golog.Fatal(err.Error())
// 	}
// 	golog.Info("-- Leaving FindAll")
// 	return res
// }

// func (e *ExperimentService) FindOne(id string) *models.ExperimentModel {

// 	golog.Info("-- Entering FindOne with id: ", id)

// 	clt := mongodb.GetDefaultMongoClientPanic()

// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		return nil
// 	}
// 	experiment_collection := clt.GetCollection("Experiment")
// 	result := experiment_collection.FindOne(context.Background(), bson.M{"_id": oid})
// 	if result == nil {
// 		golog.Fatal("Error getting document")
// 		return nil
// 	}
// 	res := &models.ExperimentModel{}
// 	result.Decode(res)
// 	golog.Info("-- Leaving FindOne")

// 	return res
// }

// func (e *ExperimentService) DeleteOne(id string) error {
// 	golog.Info("-- Entering DeleteOne with id: ", id)

// 	clt := mongodb.GetDefaultMongoClientPanic()

// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		return err
// 	}
// 	experiment_collection := clt.GetCollection("Experiment")

// 	result, err := experiment_collection.DeleteOne(context.Background(), bson.M{"_id": oid})
// 	if err != nil {
// 		golog.Error(err.Error())

// 		return err
// 	}

// 	golog.Info(result)
// 	return nil
// }
