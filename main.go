package main

import (
	"context"
	"fmt"

	"main/server"
)

func main() {
	// Create root context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create HTTP server instance to listen on all interfaces.
	address := fmt.Sprintf("%s:%s", "127.0.0.1", "8000")
	server := server.CreateServer(ctx, address)

	fmt.Printf("Initialization complete, listening on %s...\n", address)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
