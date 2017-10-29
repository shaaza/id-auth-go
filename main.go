package main

import (
	"fmt"

	"micro-auth/api"
)

func main() {
	server := api.Server()
	server.Run(fmt.Sprintf(":%d", 8080))
}
