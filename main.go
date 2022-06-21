package main

import (
	"fmt"
	_ "go_restful_api_playground/database"
	models "go_restful_api_playground/models"
	utils "go_restful_api_playground/utils"
)

// Swagger
// @title Demo GO Restful API
// @version 0.0.1
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	r := setupRouter()
	setting(r)

	migrate()
	// Demo Process
	DemoProcess()

	// Runner
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}

func DemoProcess() {
	demoUser1Password, _ := utils.HashingPassword("admin")

	_, _ = models.User.Insert(models.User{}, "admin", demoUser1Password, "admin")

	demoUser2Password, _ := utils.HashingPassword("admin2")
	_, _ = models.User.Insert(models.User{}, "admin2", demoUser2Password, "admin2")
}
