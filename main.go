package main

import (
	"context"
	"fmt"

	"main/config"
	"main/database"
	"main/server"
)

func main() {
	// Create root context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup database module.
	database.Initialize(ctx)
	defer database.Finalize()

	// Create HTTP server instance to listen on all interfaces.
	address := fmt.Sprintf("%s:%s",
		config.GetString("SERVER_LISTEN_ADDRESS"),
		config.GetString("SERVER_LISTEN_PORT"))
	server := server.CreateServer(ctx, address)

	fmt.Printf("Initialization complete, listening on %s...\n", address)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
