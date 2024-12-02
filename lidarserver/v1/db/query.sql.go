// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package lidardb

import (
	"context"
	"database/sql"
)

const createExperiment = `-- name: CreateExperiment :one
INSERT INTO Experiment (
    title, starttime, comments, wavelen, vertres, accum
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, starttime, title, comments, wavelen, vertres, accum
`

type CreateExperimentParams struct {
	Title     string
	Starttime interface{}
	Comments  string
	Wavelen   float64
	Vertres   float64
	Accum     int64
}

func (q *Queries) CreateExperiment(ctx context.Context, arg CreateExperimentParams) (Experiment, error) {
	row := q.db.QueryRowContext(ctx, createExperiment,
		arg.Title,
		arg.Starttime,
		arg.Comments,
		arg.Wavelen,
		arg.Vertres,
		arg.Accum,
	)
	var i Experiment
	err := row.Scan(
		&i.ID,
		&i.Starttime,
		&i.Title,
		&i.Comments,
		&i.Wavelen,
		&i.Vertres,
		&i.Accum,
	)
	return i, err
}

const createMeasurement = `-- name: CreateMeasurement :one
INSERT INTO Measurement(
    starttime, profcnt, proflen, reprate, prof_type, profile,  experiment_id
) VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id, starttime, profcnt, proflen, reprate, prof_type, profile, experiment_id
`

type CreateMeasurementParams struct {
	Starttime    interface{}
	Profcnt      int64
	Proflen      int64
	Reprate      int64
	ProfType     string
	Profile      interface{}
	ExperimentID sql.NullInt64
}

func (q *Queries) CreateMeasurement(ctx context.Context, arg CreateMeasurementParams) (Measurement, error) {
	row := q.db.QueryRowContext(ctx, createMeasurement,
		arg.Starttime,
		arg.Profcnt,
		arg.Proflen,
		arg.Reprate,
		arg.ProfType,
		arg.Profile,
		arg.ExperimentID,
	)
	var i Measurement
	err := row.Scan(
		&i.ID,
		&i.Starttime,
		&i.Profcnt,
		&i.Proflen,
		&i.Reprate,
		&i.ProfType,
		&i.Profile,
		&i.ExperimentID,
	)
	return i, err
}

const deleteExperimentById = `-- name: DeleteExperimentById :exec
DELETE
FROM Experiment
WHERE ID = ?
`

func (q *Queries) DeleteExperimentById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExperimentById, id)
	return err
}

const deleteMeasurementById = `-- name: DeleteMeasurementById :exec
DELETE
FROM Measurement
WHERE ID = ?
`

func (q *Queries) DeleteMeasurementById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMeasurementById, id)
	return err
}

const getAllExperiments = `-- name: GetAllExperiments :many
SELECT id, starttime, title, comments, wavelen, vertres, accum FROM Experiment ORDER BY starttime
`

func (q *Queries) GetAllExperiments(ctx context.Context) ([]Experiment, error) {
	rows, err := q.db.QueryContext(ctx, getAllExperiments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Experiment
	for rows.Next() {
		var i Experiment
		if err := rows.Scan(
			&i.ID,
			&i.Starttime,
			&i.Title,
			&i.Comments,
			&i.Wavelen,
			&i.Vertres,
			&i.Accum,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMeasurements = `-- name: GetAllMeasurements :many
SELECT id, starttime, profcnt, proflen, reprate, prof_type, profile, experiment_id FROM Measurement
`

func (q *Queries) GetAllMeasurements(ctx context.Context) ([]Measurement, error) {
	rows, err := q.db.QueryContext(ctx, getAllMeasurements)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Measurement
	for rows.Next() {
		var i Measurement
		if err := rows.Scan(
			&i.ID,
			&i.Starttime,
			&i.Profcnt,
			&i.Proflen,
			&i.Reprate,
			&i.ProfType,
			&i.Profile,
			&i.ExperimentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMeasurementsForExpId = `-- name: GetAllMeasurementsForExpId :many
SELECT id, starttime, profcnt, proflen, reprate, prof_type, profile, experiment_id FROM Measurement
WHERE experiment_id = ? ORDER BY starttime
`

func (q *Queries) GetAllMeasurementsForExpId(ctx context.Context, experimentID sql.NullInt64) ([]Measurement, error) {
	rows, err := q.db.QueryContext(ctx, getAllMeasurementsForExpId, experimentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Measurement
	for rows.Next() {
		var i Measurement
		if err := rows.Scan(
			&i.ID,
			&i.Starttime,
			&i.Profcnt,
			&i.Proflen,
			&i.Reprate,
			&i.ProfType,
			&i.Profile,
			&i.ExperimentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExperimentById = `-- name: GetExperimentById :one
SELECT id, starttime, title, comments, wavelen, vertres, accum FROM Experiment
WHERE ID=? LIMIT 1
`

func (q *Queries) GetExperimentById(ctx context.Context, id int64) (Experiment, error) {
	row := q.db.QueryRowContext(ctx, getExperimentById, id)
	var i Experiment
	err := row.Scan(
		&i.ID,
		&i.Starttime,
		&i.Title,
		&i.Comments,
		&i.Wavelen,
		&i.Vertres,
		&i.Accum,
	)
	return i, err
}

const getMeasurement = `-- name: GetMeasurement :one
SELECT id, starttime, profcnt, proflen, reprate, prof_type, profile, experiment_id FROM Measurement
WHERE ID=? LIMIT 1
`

func (q *Queries) GetMeasurement(ctx context.Context, id int64) (Measurement, error) {
	row := q.db.QueryRowContext(ctx, getMeasurement, id)
	var i Measurement
	err := row.Scan(
		&i.ID,
		&i.Starttime,
		&i.Profcnt,
		&i.Proflen,
		&i.Reprate,
		&i.ProfType,
		&i.Profile,
		&i.ExperimentID,
	)
	return i, err
}
