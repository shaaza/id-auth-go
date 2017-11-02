package main

import (
	"fmt"

	"micro-auth/api"
	"micro-auth/config"
	"micro-auth/migrations"
)

func main() {
	config.Load()

	migrations.Up(config.App.Database.Name, config.App.Database.Dialect, config.App.Database.MigrationsDir)
	server := api.Server()
	server.Run(fmt.Sprintf(":%d", config.App.Server.Port))
}
