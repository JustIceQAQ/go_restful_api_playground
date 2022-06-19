package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (User) Insert(account string, password string, username string, db *gorm.DB) (*User, error) {
	user := User{
		Account:  account,
		Password: password,
		Username: username,
	}
	if res := db.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
