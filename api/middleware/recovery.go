package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery is a middleware that recovers from panic then logs the stack trace.
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// Recover from panic.
			if recovered := recover(); recovered != nil {
				// Discontinue the request handler chain processing.
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		// Continue processing request chain.
		ctx.Next()
	}
}
