package main

import (
	"fmt"
	handler "go_restful_api_playground/handler"
	models "go_restful_api_playground/models"
	utils "go_restful_api_playground/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Swagger
// @title Demo GO Restful API
// @version 0.0.1
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

var DB *gorm.DB

func main() {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrate(db)
	handlers := handler.NewHandler(db)
	DB = db
	r := setupRouter(handlers)
	setting(r)

	// Demo Process
	DemoProcess()

	// Runner
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}

func DemoProcess() {
	demoUser1Password, _ := utils.HashingPassword("admin")

	_, _ = models.User.Insert(models.User{}, "admin", demoUser1Password, "admin", DB)

	demoUser2Password, _ := utils.HashingPassword("admin2")
	_, _ = models.User.Insert(models.User{}, "admin2", demoUser2Password, "admin2", DB)
}
