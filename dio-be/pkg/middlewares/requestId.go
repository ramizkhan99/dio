package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ramizkhan99/dio-be/pkg/logger"
)

func addRequestUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := uuid.New()
		ctx.Set("id", uuid)
		logger.GetSugaredLogger().Infof("Serving request %s started", uuid)
		ctx.Next()
		logger.GetSugaredLogger().Infof("Serving request %s finished", uuid)
	}
}
