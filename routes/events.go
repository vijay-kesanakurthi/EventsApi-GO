package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

func getEvent(ctx *gin.Context) {
	var event *models.Event
	param := ctx.Param("id")
	var id, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	event, err = models.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := ctx.GetInt("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created",
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

func updateEvent(context *gin.Context) {
	var event models.Event
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := context.GetInt("userId")

	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized user"})
		return
	}

	err = models.UpdateEventByID(id, &event)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "update event"})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cannot find event"})
		return
	}
	userId := ctx.GetInt("userId")

	if event.UserId != userId {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized user"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cannot delete event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}
