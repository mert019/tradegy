package database

import (
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/responsemodels"
)

type IOrderRepository interface {
	CreateOrder(order dbmodels.Order) (dbmodels.Order, error)
	GetAmountEffectiveOrdersByAssetID(asssetID uint, userID uint) []dbmodels.Order
	GetOpenOrdersByOrderType(orderTypeID int) []dbmodels.Order
	InsertUpdateOrder(order dbmodels.Order) (dbmodels.Order, error)
	GetExecutedOrdersByUserId(userId uint) []dbmodels.Order
	GetSellableAssetsByUserId(userId uint64) []responsemodels.SellAssetsResponse
	GetOrdersByUserId(userId uint) []dbmodels.Order
}
