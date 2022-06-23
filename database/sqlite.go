package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func init() {

	var err error
	dbDns := os.Getenv("DB_DNS")
	Db, err = gorm.Open(sqlite.Open(dbDns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(err)
	}

}
