package database

import (
	"fmt"
	Config "go_restful_api_playground/configs"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	C := Config.Cfg
	if "debug" == C.GinMode {
		Db, err = gorm.Open(sqlite.Open(C.DatabaseUrl), &gorm.Config{})
		fmt.Println("💬 database Using [sqlite]")

	} else if "release" == C.GinMode {
		Db, err = gorm.Open(postgres.Open(C.DatabaseUrl), &gorm.Config{})
		fmt.Println("💬 database Using [postgres]")
	}

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(err)
	}

}
