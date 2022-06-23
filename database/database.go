package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func init() {

	var err error
	databaseUrl := os.Getenv("DATABASE_URL")
	Db, err = gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(err)
	}

}
