package main

import (
	"fmt"

	"github.com/urfave/cli"
	"micro-auth/api"
	"micro-auth/config"
	"micro-auth/migrations"
	"os"
	"runtime"
)

func main() {
	config.Load()

	server := api.Server()

	runtime.GOMAXPROCS(runtime.NumCPU())

	clientApp := cli.NewApp()
	clientApp.Name = "Maps API"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start Http Server",
			Action: func(c *cli.Context) error {
				server.Run(fmt.Sprintf(":%d", config.App.Server.Port))
				return nil
			},
		},
		{
			Name:        "db:migrate:up",
			Description: "Start Http Server",
			Action: func(c *cli.Context) error {
				migrations.Up(config.App.Database.Name, config.App.Database.Dialect, config.App.Database.MigrationsDir)
				return nil
			},
		},
		{
			Name:        "db:migrate:down",
			Description: "Start Http Server",
			Action: func(c *cli.Context) error {
				migrations.Down(config.App.Database.Name, config.App.Database.Dialect, config.App.Database.MigrationsDir)
				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}

}
