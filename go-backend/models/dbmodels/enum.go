package models

import "gorm.io/gorm"

type Enum struct {
	gorm.Model
	Area string `json:"area" gorm:"column:area"`
	Name string `json:"name" gorm:"column:name"`
	Code int64  `json:"code" gorm:"column:code;unique"`
}
