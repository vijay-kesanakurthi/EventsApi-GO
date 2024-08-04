package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/db"
	"rest-api/models"
)

func main() {
	server := gin.Default()

	db.InitDB()
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
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create event",
	})
}
func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)
}
