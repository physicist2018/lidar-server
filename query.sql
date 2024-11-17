-- name: GetExperiment :one
SELECT * FROM Experiment
WHERE ID=? LIMIT 1;

-- name: CreateExperiment :one
INSERT INTO Experiment (
    title, starttime, comments, wavelen, vertres, accum
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

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