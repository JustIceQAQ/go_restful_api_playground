package models

import (
	orm "go_restful_api_playground/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account             string              `gorm:"uniqueIndex" json:"account" binding:"required"`
	Password            string              `json:"password" binding:"required"`
	Username            string              `json:"username"`
	PersonalInformation PersonalInformation `gorm:"foreignKey:UserId" json:"personal_information"`
	File                []File              `gorm:"foreignKey:UserId" json:"files"`
}

type PersonalInformation struct {
	gorm.Model
	UserId uint  `json:"userid"`
	Age    uint8 `json:"age"`
	Sex    uint8 `json:"sex"`
}

func (User) Insert(account string, password string, username string) (*User, error) {
	user := User{
		Account:             account,
		Password:            password,
		Username:            username,
		PersonalInformation: PersonalInformation{Age: 10, Sex: 1},
	}

	if res := orm.Db.Create(&user); res.Error != nil {
		return nil, nil
	}
	return &user, nil
}
