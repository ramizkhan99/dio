package models

import (
	"github.com/gin-gonic/gin"
)

const (
	USER_NOT_FOUND = "UserNotFoundError"
)

type Error struct {
	statusCode  int
	errorType   string
	title       string
	description string
}

func SendError(ctx *gin.Context, statusCode int, title, description, errorType string) {
	err := Error{
		statusCode:  statusCode,
		title:       title,
		errorType:   errorType,
		description: description,
	}
	ctx.JSON(statusCode, gin.H{"error": err})
}
