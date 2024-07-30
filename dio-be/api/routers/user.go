package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/ramizkhan99/dio-be/internal/user"
	"github.com/ramizkhan99/dio-be/pkg/logger"
)

type UserController interface {
}

// All routes mentioned here are prefixed by /users/
func registerUserRoutes(rg *gin.RouterGroup) {
	rg.POST("", createUser)
	rg.GET("/:id", getUserByID)
	rg.PATCH("/:id", updateUserbyID)
	rg.GET("/user/:username", getUserByUsername)
}

type UpdateUsernamePayload struct {
	NewUsername string `json:"username"`
}

func createUser(ctx *gin.Context) {
	user, err := userService.NewUser(ctx)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to create user %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func getUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := userService.GetUserByID(id)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to get user :: %+v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func getUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := userService.GetUserByUsername(username)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to get user :: %+v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func updateUserbyID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := userService.GetUserByID(id)
	if err != nil {
		logger.GetSugaredLogger().Errorf("failed to get user :: %+v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var payload UpdateUsernamePayload
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		logger.GetSugaredLogger().Errorf("unable to parse username update payload :: %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.UpdateUsername(payload.NewUsername)
	if err != nil {
		logger.GetSugaredLogger().Errorf("unable to update username :: %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
