package server

import (
	"context"
	"net/http"

	"main/api"
)

// CreateServer creates an HTTP server listening on the specified address.
func CreateServer(ctx context.Context, address string) *http.Server {
	// Setup HTTP Server.
	server := &http.Server{
		Addr:    address,
		Handler: api.GetRouter(),
	}

	return server
}
