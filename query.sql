-- name: GetExperimentById :one
SELECT * FROM Experiment
WHERE ID=? LIMIT 1;

-- name: CreateExperiment :one
INSERT INTO Experiment (
    title, starttime, comments, wavelen, vertres, accum
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteExperimentById :exec
DELETE
FROM Experiment
WHERE ID = ?;

-- name: GetAllExperiments :many
SELECT * FROM Experiment;

-- name: GetMeasurement :one
SELECT * FROM Measurement
WHERE ID=? LIMIT 1;

-- name: GetAllMeasurements :many
SELECT * FROM Measurement;

-- name: GetAllMeasurementsForExpId :many
SELECT * FROM Measurement
WHERE experiment_id = ?;

-- name: CreateMeasurement :one
INSERT INTO Measurement(
    starttime, profcnt, proflen, reprate, profile,  experiment_id
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteMeasurementById :exec
DELETE
FROM Measurement
WHERE ID = ?;