package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// The global HTTP router instance and root group.
var router *gin.Engine
var root *gin.RouterGroup

// CreateServer creates an HTTP server listening on the specified address.
func CreateServer(ctx context.Context, address string) *http.Server {
	router = gin.New()

	// Setup HTTP Server.
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	// Setup endpoint
	root = router.Group("")
	root.GET("", GetIndex)

	return server
}

func GetIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "index",
	})
}
