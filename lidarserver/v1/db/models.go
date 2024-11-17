// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package lidardb

import (
	"database/sql"
)

type Experiment struct {
	ID        int64
	Starttime interface{}
	Title     string
	Comments  string
	Wavelen   float64
	Vertres   float64
	Accum     int64
}

type Measurement struct {
	ID           int64
	Starttime    interface{}
	Profcnt      int64
	Proflen      int64
	Reprate      int64
	Profile      interface{}
	ExperimentID sql.NullInt64
}
