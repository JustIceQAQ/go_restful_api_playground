package main

import "gorm.io/gorm"
import Models "go_restful_api_playground/models"

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&Models.User{}); err != nil {
		panic(err)
	}
}
