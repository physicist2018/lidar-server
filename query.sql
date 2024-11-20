-- name: GetExperimentById :one
SELECT * FROM Experiment
WHERE ID=? LIMIT 1;

-- name: CreateExperiment :one
INSERT INTO Experiment (
    Title, Start_Time, Comments, Wavelen, Vert_Res, Accum
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteExperimentById :exec
DELETE
FROM Experiment
WHERE ID = ?;

-- name: GetAllExperiments :many
SELECT * FROM Experiment ORDER BY Start_Time;

-- name: GetMeasurement :one
SELECT * FROM Measurement
WHERE ID=? LIMIT 1;

-- name: GetAllMeasurements :many
SELECT * FROM Measurement;

-- name: GetAllMeasurementsForExpId :many
SELECT * FROM Measurement
WHERE Experiment_ID = ? ORDER BY Start_Time;

-- name: CreateMeasurement :one
INSERT INTO Measurement(
    Start_Time, Prof_Cnt, Prof_Len, Rep_Rate, Profile_Data,  Experiment_ID
) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteMeasurementById :exec
DELETE
FROM Measurement
WHERE ID = ?;
