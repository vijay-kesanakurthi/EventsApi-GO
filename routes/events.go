package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

// @Summary Get event by id
// @Tags Events
// @Param id path int true "event id"
// @Router /events/{id} [get]
// @Success      200 {object} models.Event
// @Failure    403  {object}  error
// @Failure    404  {object}  error
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

// @Summary Create event
// @Tags Events
// @Security bearerAuth
// @Param Authorization header string true "Bearer token"
// @Param event body models.EventModel true "event"
// @Router /events/ [post]
// @Success      201 {object} string
// @Failure     400  {object}  error
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

// @Summary Get all events
// @Tags Events
// @Router /events/ [get]
// @Success      200 {object} []models.Event
func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)
}

// @Summary Update event
// @Tags Events
// @Security bearerAuth
// @Param Authorization header string true "Bearer token"
// @Param id path int true "event id"
// @Param event body models.EventModel true "event"
// @Router /events/{id} [put]
// @Success      200 {object} []models.Event
// @Failure     403  {object} error
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

// @Summary Delete event
// @Tags Events
// @Security bearerAuth
// @Param Authorization header string true "Bearer token"
// @Param id path int true "event id"
// @Router /events/{id} [delete]
// @Success      200 {object} string
// @Failure     403  {object} error
// @Failure     404  {object} error
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
