package models

import (
	orm "go_restful_api_playground/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account             string              `json:"account"`
	Password            string              `json:"password"`
	Username            string              `json:"username"`
	PersonalInformation PersonalInformation `gorm:"foreignKey:ID"`
}

type PersonalInformation struct {
	gorm.Model
	Age uint8 `json:"age"`
	Sex uint8 `json:"sex"`
}

func (User) Insert(account string, password string, username string) (*User, error) {
	user := User{
		Account:             account,
		Password:            password,
		Username:            username,
		PersonalInformation: PersonalInformation{Age: 10, Sex: 1},
	}

	if res := orm.Db.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
