package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	auth := server.Group("/")

	auth.Use(middlewares.AuthMiddleware)
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
