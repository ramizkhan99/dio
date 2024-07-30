package middlewares

import "github.com/ramizkhan99/dio-be/api/server"

func RegisterMiddlewares() {
	server.GetServer().Use(addCORSMiddleware())
	server.GetServer().Use(addRequestUUID())
}
