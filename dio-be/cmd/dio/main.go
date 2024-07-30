package main

import (
	router "github.com/ramizkhan99/dio-be/api/routers"
	"github.com/ramizkhan99/dio-be/api/server"
	"github.com/ramizkhan99/dio-be/internal/config"
	"github.com/ramizkhan99/dio-be/internal/db"
	"github.com/ramizkhan99/dio-be/pkg/logger"
	"github.com/ramizkhan99/dio-be/pkg/middlewares"
)

func main() {
	config.LoadConfig()
	logger.InitializeLogger()

	server.CreateServer()
	middlewares.RegisterMiddlewares()
	router.RegisterRoutes()

	server.StartServer()
	defer db.Disconnect()
}
