package api

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// Setup domains router group.
	root := GetRoot().Group("users")
	root.GET("", GetUsers)
	root.GET("/:id", GetUser)
}

// GetUsers ...
func GetUsers(ctx *gin.Context) {
	//Mock reading user data from DB
	users, err := func() ([]string, error) {
		// DB query latency
		time.Sleep(time.Millisecond)
		// Mock DB query with 1/10 chance failed
		if rand.Intn(100) < 10 {
			return nil, errors.New("DB query failed")
		}
		return []string{"User1", "User2", "User3"}, nil
	}()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Set results to context.
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUser(ctx *gin.Context) {
	user_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	//Mock reading user data from DB
	user, err := func(user_id int) (string, error) {
		// Mock DB query latency
		time.Sleep(time.Millisecond)

		if user_id <= 0 || user_id > 3 {
			return "", errors.New("user not found")
		}

		user := fmt.Sprintf("User%d", user_id)
		return user, nil
	}(user_id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Set results to context.
	ctx.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}
