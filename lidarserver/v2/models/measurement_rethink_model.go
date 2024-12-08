package models

import "time"

type MeasurementModel struct {
	ID             string    `rethinkdb:"id,omitempty" json:"id,omitempty"`
	StartTime      time.Time `rethinkdb:"start_time" json:"start_time"`
	RepetitionRate int       `rethinkdb:"repetition_rate" json:"repetition_rate"` // in Herz
	Wavelength     float64   `rethinkdb:"wavelength" json:"wavelength"`           // in nm
	ProfileLength  int       `rethinkdb:"profile_length" json:"profile_length"`
	ProfileCount   int       `rethinkdb:"profile_count" json:"profile_count"`
	Data           []float64 `bsrethinkdbon:"data" json:"data"`
}
