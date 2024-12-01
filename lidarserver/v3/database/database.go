package database

import (
	"github.com/kataras/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lidarserver.sqlc/app/lidarserver/v3/model"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		golog.Fatal("Canntot connect to DB...")
		panic(err)
	}

	golog.Info("Connected to DB...")
}

func Migrate() {
	Instance.AutoMigrate(&model.Experiment{}, &model.Measurement{}, &model.Result{})
	golog.Info("DB Migration completed ...")
}
