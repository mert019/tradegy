package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
	Code   string `json:"code" gorm:"column:code"`
	ApiId  string `json:"api_id" gorm:"column:api_id"`
	TypeId int64  `json:"type_id" gorm:"column:type_id"`

	Type Enum `json:"type" gorm:"foreignKey:TypeId;references:Code"`
}
