package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey string = "SuperSecretKey"

func GenerateToken(username string, userId int) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"user_id":  userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return token, err
	}
	return token, nil
}

func VerfyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return jwtToken, errors.New("parse jwt token fail")
	}
	if !jwtToken.Valid {
		return jwtToken, errors.New("token is invalid")
	}
	return jwtToken, nil
}
