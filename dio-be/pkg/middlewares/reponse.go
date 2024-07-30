package middlewares

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func printResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		json.MarshalIndent(ctx.Request.Response, ",", "\t")
		ctx.Next()
	}
}
