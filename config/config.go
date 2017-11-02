package config

import (
	"github.com/spf13/viper"
)

type AppServer struct {
	Port int
}

type Database struct {
	Name          string
	MigrationsDir string
	Dialect       string
}
type Config struct {
	Server   AppServer
	Database Database
}

var App Config
var DB Database

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
		Database: Database{
			Name:          getStringOrPanic("DATABASE_NAME"),
			MigrationsDir: getStringOrPanic("DATABASE_MIGRATIONS_DIR"),
			Dialect:       getStringOrPanic("DATABASE_DIALECT"),
		},
	}
}
