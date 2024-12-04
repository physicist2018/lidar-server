-- name: CreateTableExperiment :exec
-- создает таблицу эксперимента в БД
CREATE TABLE Experiment(
    ID integer primary key autoincrement,
    Start_Time timestamp not null default now,
    Title varchar(100) not null default 'no-title',
    Comments varchar(500) not null default 'no-comments',
    Wavelen real not null default 532.0,
    Vert_Res real not null default 1500.0,
    Accum integer not null default 10,
    unique (Start_Time, Title, Accum)
);


-- name: CreateTableMeasurement :exec
-- создает таблицу измерений
CREATE TABLE Measurement(
    ID integer primary key autoincrement,
    Start_Time timestamp not null default now,
    Prof_Cnt integer not null default 1,
    Prof_Len integer not null default 512,
    Rep_Rate integer not null default 10,
    Profile_Data json,
    Experiment_ID integer,
    CONSTRAINT experiment_fk FOREIGN KEY (Experiment_ID) REFERENCES Experiment(ID)
);


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
