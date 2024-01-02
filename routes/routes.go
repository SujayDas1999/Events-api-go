package routes

import (
	"events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/api/events", GetEvents)
	server.GET("/api/events/:id", GetEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/api/events", CreateEvent)
	authenticated.PUT("/api/events/:id", UpdateEvent)
	authenticated.DELETE("/api/events/:id", DeleteEvent)
	authenticated.GET("/api/users", GetAll)
	authenticated.POST("/api/register/:id", CreateRegistration)
	authenticated.DELETE("/api/cancel/:id", CancelRegistration)

	server.POST("/api/users/signup", CreateUser)
	server.POST("/api/users/login", Login)

}
