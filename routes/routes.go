package routes

import "github.com/gin-gonic/gin"

func EventRoutes(server *gin.Engine) {
	server.GET("/api/events", GetEvents)
	server.GET("/api/events/:id", GetEventById)
	server.POST("/api/events", CreateEvent)
	server.PUT("/api/events/:id", UpdateEvent)
}
