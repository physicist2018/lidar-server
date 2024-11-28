package db

import (
	"time"
)

// Measurement - Description of a single data profile
type Measurement struct {
	Id         string    `rethinkdb:"id,omitempty"`
	StartTime  time.Time `rethinkdb:"start_time"`
	StopTime   time.Time `rethinkdb:"stop_time"`
	ProfCnt    int       `rethinkdb:"prof_cnt"`
	ProfLen    int       `rethinkdb:"prof_len"`
	Wavelen    float64   `rethinkdb:"wavelen"`
	RepRate    int       `rethinkdb:"rep_rate"`
	ProfileDat []float64 `rethinkdb:"profile_dat"`
	ProfileDak []float64 `rethinkdb:"profile_dak"`
}

// Experiment - Description of a single lidar sounding experiment
type Experiment struct {
	Id             string         `rethinkdb:"id,omitempty"`
	Title          string         `rethinkdb:"title"`
	Comments       string         `rethinkdb:"comments"`
	ExperimentTime time.Time      `rethinkdb:"experiment_time,omitempty"`
	VertRes        float64        `rethinkdb:"vert_res"`
	Accum          int            `rethinkdb:"accum"`
	Data           []*Measurement `rethinkdb:"measurements"`
}

//,reference" rethinkdb_ref:"id"
