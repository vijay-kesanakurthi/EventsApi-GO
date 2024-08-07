package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

func registerEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.GetInt("userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is no such event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "canot register event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully registered"})
}

func deleteRegistration(ctx *gin.Context) {
	regId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url"})
	}
	reg, err := models.FindRegistrationById(regId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete registration"})
		return
	}
	userId := ctx.GetInt("userId")

	if reg.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized to delete registration"})
		return
	}
	err = reg.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete registration"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted registration"})
}

func getRegistrations(ctx *gin.Context) {
	allRegistrations, err := models.FindAllRegistrations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get registrations", "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"registrations": allRegistrations})
}
