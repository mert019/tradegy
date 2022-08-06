package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
