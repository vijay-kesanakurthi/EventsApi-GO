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
	auth.POST("/events/:id/register", registerEvent)
	auth.DELETE("/events/:id/deregister", deleteRegistration)

	auth.GET("/registrations", getRegistrations)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
