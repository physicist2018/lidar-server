package model

import (
	"time"

	"gorm.io/gorm"
)

type Measurement struct {
	gorm.Model
	StartTime    time.Time `gorm:"type:datetime(0);default:now()"`
	StopTime     time.Time `gorm:"type:datetime(0);default:now()"`
	ProfLen      int       `gorm:"default:512"`
	ProfCnt      int       `gorm:"default:100"`
	RepRate      int       `gorm:"default:10"`
	MesType      int       `gorm:"default:1"`
	Data         []float64 `gorm:"serializer:json"`
	ExperimentID int
	ResultID     int
	Result       Result
}
