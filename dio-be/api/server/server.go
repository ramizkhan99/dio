package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ramizkhan99/dio-be/internal/config"
)

// TODO: Use server struct here

var server *gin.Engine

func CreateServer() *gin.Engine {
	server = gin.New()
	return server
}

func StartServer() {
	server.Run(fmt.Sprintf("%s:%s", config.GetUrl(), config.GetPort()))
}

func GetServer() *gin.Engine {
	return server
}
