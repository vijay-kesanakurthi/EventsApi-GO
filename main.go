package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rest-api/db"
	"rest-api/docs"
	"rest-api/routes"
)

func main() {
	server := gin.Default()

	db.InitDB()

	docs.SwaggerInfo.Title = "Event Registration API"
	docs.SwaggerInfo.Description = "This is a server for registering users created events."
	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
