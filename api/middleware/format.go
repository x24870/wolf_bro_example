package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FormatResponse ...
func FormatResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Onto the next handler if we're not final.
		ctx.Next()

		// Get prepared response from context.
		response, exists := ctx.Get("response")
		if !exists {
			errmsg := "no response"
			ctx.String(http.StatusInternalServerError, errmsg)
			return
		}

		// Respond JSON format.
		ctx.JSON(http.StatusOK, response)
	}
}
