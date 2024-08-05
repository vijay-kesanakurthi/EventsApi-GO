package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/util"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
		return
	}
	token, err := util.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user  successfully logged in", "token": token})
}
