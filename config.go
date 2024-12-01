package main

import (
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *Config

func LoadAppConfig() {
	golog.Info("Loading server configurations ...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		golog.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		golog.Fatal(err)
	}
}
