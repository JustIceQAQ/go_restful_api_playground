package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserId   uint   `json:"userid"`
	FileName string `json:"filename"`
	Uri      string `json:"uri"`
}
