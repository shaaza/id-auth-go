package config

import (
	"github.com/spf13/viper"
)

type AppServer struct {
	Port int
}
type Config struct {
	Server AppServer
}

var App Config

func Load() {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	App = Config{
		Server: AppServer{
			Port: getIntOrPanic("APP_PORT"),
		},
	}
}
