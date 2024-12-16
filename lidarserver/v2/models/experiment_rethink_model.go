package models

import "time"

// ExperimentModel represents an experiment in the system.
// It contains fields for the experiment's ID, title, comments, start time, spatial resolution, accumulation time, data, data acquisition, and molecular data.
// The fields are tagged with the "rethinkdb" and "json" tags to specify the field names in the database and JSON responses.
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
