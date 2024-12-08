package models

import "time"

type ExperimentModel struct {
	ID            string             `rethinkdb:"id,omitempty" json:"id"`
	Title         string             `rethinkdb:"title" json:"title"`
	Comments      string             `rethinkdb:"comments" json:"comments"`
	StartTime     time.Time          `rethinkdb:"start_time" json:"start_time"`
	SpatialRes    float64            `rethinkdb:"spatial_res" json:"spatial_res"`
	AccumTime     int                `rethinkdb:"accum_time" json:"accum_time"`
	Dat           []MeasurementModel `rethinkdb:"dat" json:"dat"`
	Dak           []MeasurementModel `rethinkdb:"dak" json:"dak"`
	MolecularData MolecularModel     `rethinkdb:"molecular_data" json:"molecular_data"`
}
