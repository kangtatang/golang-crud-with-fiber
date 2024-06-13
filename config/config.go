package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port int
	}
	Database struct {
		User     string
		Password string
		Dbname   string
		Host     string
		Port     int
	}
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
