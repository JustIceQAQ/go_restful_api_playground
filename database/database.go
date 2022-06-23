package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func init() {
	var err error
	databaseUrl := os.Getenv("DATABASE_URL")
	ginMode := os.Getenv("GIN_MODE")
	if "" == ginMode {
		ginMode = "debug"
	}

	if "debug" == ginMode {
		Db, err = gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{})
		fmt.Println("💬 database Using [sqlite]")

	} else if "release" == ginMode {
		Db, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
		fmt.Println("💬 database Using [postgres]")
	}

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(err)
	}

}
