package middleware

import (
	"fmt"
	"main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authToken string

func init() {
	authToken = config.GetString("AUTH_TOKEN")
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtain authorization in header
		token := ctx.Request.Header.Get("Authorization")

		// Validate the token
		if token != authToken {
			err := fmt.Errorf("invalid token")
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.Next()
	}
}
