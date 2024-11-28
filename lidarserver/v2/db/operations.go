package db

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

const (
	EXPERIMENT_TBL  = "Experiment"
	MEASUREMENT_TBL = "Measurement"
	LIDAR_DB        = "Lidar"
)

// CreateExperiment - Creates an entry in the Experiment table
func (db *DB) CreateExperiment(e *Experiment) error {
	_, err := r.Table(EXPERIMENT_TBL).Insert(e).RunWrite(db.Session)
	if err != nil {
		return err
	}
	return nil
}

// CreateExperimentList - Creates entries in the Experiment table
func (db *DB) CreateExperimentList(e []*Experiment) error {
	_, err := r.Table(EXPERIMENT_TBL).Insert(e).RunWrite(db.Session)
	if err != nil {
		return err
	}
	return nil
}

// CreateMeasurement - Creates an entry in the Measurement table
func (db *DB) CreateMeasurement(m *Measurement) error {
	_, err := r.Table(MEASUREMENT_TBL).Insert(m).RunWrite(db.Session)
	if err != nil {
		return err
	}
	return nil
}

// CreateMeasurementList - Creates enties in the Measurement table
func (db *DB) CreateMeasurementList(m []*Measurement) error {
	_, err := r.Table(MEASUREMENT_TBL).Insert(m).RunWrite(db.Session)
	if err != nil {
		return err
	}
	return nil
}
