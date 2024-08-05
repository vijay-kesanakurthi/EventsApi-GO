package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"rest-api/util"
)

func AuthMiddleware(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := util.VerifyToken(authHeader)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	userId := claims["user_id"].(float64)
	fmt.Println("userId", claims)

	context.Set("userId", int(userId))

	context.Next()
}
