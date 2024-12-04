package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExperimentModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title"`
	Comments   string             `bson:"comments"`
	StartTime  time.Time          `bson:"start_time"`
	SpatialRes float64            `bson:"spatial_res"`
	AccumTime  int                `bson:"accum_time"`
}
