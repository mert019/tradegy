package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	BuyAmount         float64   `json:"buy_amount" gorm:"column:buy_amount"`
	SellAmount        float64   `json:"sell_amount" gorm:"column:sell_amount"`
	Limit             float64   `json:"limit" gorm:"column:limit"`
	UserID            uint      `json:"user_id" gorm:"column:user_id"`
	OrderTypeID       int64     `json:"order_type_id" gorm:"column:order_type_id"`
	BuyAssetID        uint      `json:"buy_asset_id" gorm:"column:buy_asset_id"`
	SellAssetID       uint      `json:"sell_asset_id" gorm:"column:sell_asset_id"`
	OrderStatusID     int64     `json:"order_status_id" gorm:"column:order_status_id"`
	ExecutionDateTime time.Time `json:"execution_date_time" gorm:"column:execution_date_time"`

	User        User
	BuyAsset    Asset `gorm:"foreignKey:BuyAssetID"`
	SellAsset   Asset `gorm:"foreignKey:SellAssetID"`
	OrderType   Enum  `gorm:"foreignKey:OrderTypeID;references:Code"`
	OrderStatus Enum  `gorm:"foreignKey:OrderStatusID;references:Code"`
}
