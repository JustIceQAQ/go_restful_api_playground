package main

import (
	"fmt"
	handler "go_restful_api_playground/handler"
	models "go_restful_api_playground/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	_, _ = models.User.Insert(models.User{}, "admin", "admin", "admin", DB)
	_, _ = models.User.Insert(models.User{}, "admin2", "admin2", "admin2", DB)

	// Runner
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
