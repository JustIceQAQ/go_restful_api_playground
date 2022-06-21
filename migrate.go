package main

import (
	orm "go_restful_api_playground/database"
	Models "go_restful_api_playground/models"
)

func migrate() {
	if err := orm.Db.AutoMigrate(&Models.User{}); err != nil {
		panic(err)
	}
	if err := orm.Db.AutoMigrate(&Models.PersonalInformation{}); err != nil {
		panic(err)
	}
}
