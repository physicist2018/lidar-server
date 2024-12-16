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

// FindAll is a function that retrieves all experiments from the database.
// It connects to the database, runs a query to fetch all experiments, and returns them as a slice of ExperimentModel.
// If there is an error during the database operation, it logs the error and returns nil.
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

// FindOne is a function that retrieves a single experiment from the database by its ID.
// It connects to the database, runs a query to fetch the experiment with the specified ID, and returns it as an ExperimentModel.
// If there is an error during the database operation, it logs the error and returns nil.
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

// UpdateOne is a function that updates a single experiment in the database by its ID.
// It connects to the database, runs a query to update the experiment with the specified ID, and returns the result as a ResponseModel.
// If there is an error during the database operation, it logs the error and returns nil.
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

// DeleteOne is a function that deletes a single experiment from the database by its ID.
// It connects to the database, runs a query to delete the experiment with the specified ID, and returns the result as a ResponseModel.
// If there is an error during the database operation, it logs the error and returns nil.
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

// CreateOne is a function that creates a new experiment in the database.
// It connects to the database, prepares the experiment data by generating unique IDs for its profiles and data, and inserts the experiment into the database.
// It returns the result of the insertion as a ResponseModel.
// If there is an error during the database operation, it logs the error and returns nil.
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
