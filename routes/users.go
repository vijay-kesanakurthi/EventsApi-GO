package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/util"
)

// @Summary Create user
// @Tags Users
// @Param user body models.UserModel true "User details"
// @Router /signup [post]
// @Success 201 {string} string "User created successfully"
// @Failure 400 {object} error "Bad Request"
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

// @Summary Login
// @Tags Users
// @Param user body models.UserModel true "user"
// @Router /login [post]
// @Success      200 {object} string
// @Failure     400  {object}  error
func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.Validate(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
		return
	}
	fmt.Println(user)
	token, err := util.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user  successfully logged in", "token": token})
}
