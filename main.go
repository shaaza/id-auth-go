package main

import (
	"fmt"

	"micro-auth/api"
	"micro-auth/config"
)

func main() {
	config.Load()
	server := api.Server()
	server.Run(fmt.Sprintf(":%d", config.App.Server.Port))
}
