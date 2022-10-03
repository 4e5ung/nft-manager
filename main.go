package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"manage-template/config"
	"manage-template/docs"
	"manage-template/domain/token/controller"
	"os"
)

var testNetworkInfo = []config.NetworkInfo{config.EdgeTestnet}

// @title Swagger Example API
// @version 1.0
// @description This is a sample mange-template server.
// @BasePath /api/v1
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Manage-Template API"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("api/v1")
	controller.SetRouterToken(v1)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
