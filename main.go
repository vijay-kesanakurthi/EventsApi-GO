package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rest-api/db"
	"rest-api/routes"
)

func main() {
	server := gin.Default()

	db.InitDB()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
