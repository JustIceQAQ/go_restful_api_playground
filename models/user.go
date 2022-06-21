package models

import (
	orm "go_restful_api_playground/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (User) Insert(account string, password string, username string) (*User, error) {
	user := User{
		Account:  account,
		Password: password,
		Username: username,
	}

	if res := orm.Db.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
