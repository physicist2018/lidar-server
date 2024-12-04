package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasurementModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	StartTime      time.Time          `bson:"start_time"`
	RepetitionRate int                `bson:"repetition_rate"` // in Herz
	Wavelength     float64            `bson:"wavelength"`      // in nm
	ProfileLength  int                `bson:"profile_length"`
	ProfileCount   int                `bson:"profile_count"`
	Data           []float64          `bson:"data"`
	Experiment     primitive.ObjectID `bson:"experiment"`
}
