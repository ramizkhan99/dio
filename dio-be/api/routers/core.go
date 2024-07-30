package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramizkhan99/dio-be/pkg/logger"
)

type CoreResponse struct {
}

func registerCoreRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", ping)
}

func ping(ctx *gin.Context) {
	logger.GetLogger().Info("Ping called")
	ctx.String(http.StatusOK, "This is working!!!")
}
