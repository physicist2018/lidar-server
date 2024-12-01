package model

import (
	"time"

	"gorm.io/gorm"
)

type Experiment struct {
	gorm.Model

	Title        string        `gorm:"not null;default:'no-title'"`
	Comments     string        `gorm:"not null;default:'no-comments'"`
	ExperimentDT time.Time     `gorm:"type:datetime(0);not null;default:now()"`
	SpatStep     float64       `gorm:"not null;default:1500.0"`
	AccumTime    int           `gorm:"not null;default:1"`
	Profile      []Measurement `gorm:"foreignKey:ExperimentID"`
}
