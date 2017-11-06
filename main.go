package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
	"log"
	"micro-auth/api"
	"micro-auth/config"
	"micro-auth/migrations"
	"os"
	"runtime"

	"micro-auth/db"
)

func main() {
	config.Load()
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", config.App.Database.Name)
	databaseInstance, err := sql.Open(config.App.Database.Dialect, dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer databaseInstance.Close()
	database := db.DB{
		Instance: databaseInstance,
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	clientApp := cli.NewApp()
	clientApp.Name = "Maps API"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start Http Server",
			Action: func(c *cli.Context) error {
				api.StartServer(api.Server(database), config.App.Server)
				return nil
			},
		},
		{
			Name:        "db:migrate:up",
			Description: "Create migrations",
			Action: func(c *cli.Context) error {
				migrations.Up(config.App.Database.Name, config.App.Database.Dialect, config.App.Database.MigrationsDir)
				return nil
			},
		},
		{
			Name:        "db:migrate:down",
			Description: "Destroy migrations",
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
