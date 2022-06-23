package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_restful_api_playground/docs"
	"os"
)

func setting(app *gin.Engine) {

	ginMode := os.Getenv("GIN_MODE")

	if "" == ginMode {
		ginMode = "debug"
	}
	err := godotenv.Load(".env." + ginMode)
	if err != nil {
	}

	fmt.Println("💬 Now API is", ginMode, "Mode")

	docs.SwaggerInfo.BasePath = "/api/v1"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS
	app.Use(cors.Default())

	// DB migrate
	migrate()

	// static path
	app.Static("/media", "./media")

}
