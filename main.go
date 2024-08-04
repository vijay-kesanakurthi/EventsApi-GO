package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create event",
	})
}
func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(200, events)
}
