package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukesampson/figlet/figletlib"
)

func init() {
	// Setup domains router group.
	root := GetRoot().Group("banner")
	root.GET("/:text", GetBanner)
}

// GetBanner ...
func GetBanner(ctx *gin.Context) {
	text := ctx.Param("text")

	f, err := figletlib.GetFontByName("/fonts", "slant")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	figletlib.FPrintMsg(ctx.Writer, text, f, 80, f.Settings(), "left")
}
