package database

import dbmodels "go-backend/models/dbmodels"

type IOrderRepository interface {
	CreateOrder(order dbmodels.Order) (dbmodels.Order, error)
	GetAmountEffectiveOrdersByAssetID(asssetID uint, userID uint) []dbmodels.Order
	GetOpenOrdersByOrderType(orderTypeID int) []dbmodels.Order
	InsertUpdateOrder(order dbmodels.Order) (dbmodels.Order, error)
}
