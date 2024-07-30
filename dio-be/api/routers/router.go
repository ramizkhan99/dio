package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramizkhan99/dio-be/api/server"
	"github.com/ramizkhan99/dio-be/pkg/logger"
)

type Body struct {
	Name string `json:"name"`
}

func RegisterRoutes() {
	server.GetServer().NoRoute(handleNoRoute)
	registerCoreRoutes(server.GetServer().Group("/"))
	registerUserRoutes(server.GetServer().Group("/users"))
}

func handleNoRoute(ctx *gin.Context) {
	logger.GetSugaredLogger().Errorf("No active handler found for route %s", ctx.Request.RequestURI)
	ctx.String(http.StatusNotFound, "No page found at "+ctx.Request.RequestURI)
}
