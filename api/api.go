package api

import (
	"sync"

	"main/api/middleware"

	"github.com/gin-gonic/gin"
)

// The global HTTP router instance and root group.
var router *gin.Engine
var root *gin.RouterGroup
var once sync.Once

// GetRouter returns the global HTTP router instance.
func GetRouter() *gin.Engine {
	// Initialize API singleton instances.
	once.Do(initializeSingletons)
	return router
}

// GetRoot returns the router root group.
func GetRoot() *gin.RouterGroup {
	// Initialize API singleton instances.
	once.Do(initializeSingletons)
	return root
}

// initializeSingletons is the function called by sync.Once to intialize the
// HTTP engine and router group singleton instances.
func initializeSingletons() {
	router, root = createRouterAndGroup("")
}

// Create a clean router and a root group with the given microservice prefix.
func createRouterAndGroup(prefix string) (*gin.Engine, *gin.RouterGroup) {
	// Create a clean HTTP router engine.
	engine := gin.New()

	// Configure HTTP router engine settings.
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = false
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true

	// Create from the engine a router group with the given prefix.
	group := engine.Group(prefix)

	// Install common middleware to the router group.
	installCommonMiddleware(group)

	return engine, group
}

// installCommonMiddleware installs common middleware to the router group.
func installCommonMiddleware(group *gin.RouterGroup) {
	// Install recovery middleware, a middleware to recover & log panics.
	// NOTE: The recovery middleware should always be the last one installed.
	group.Use(middleware.Recovery())
}
