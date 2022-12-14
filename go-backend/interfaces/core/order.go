package core

import (
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/requestmodels"
	"go-backend/models/responsemodels"
)

type IOrderManager interface {
	CreateOrder(order requestmodels.CreateOrderRequest, username string) (dbmodels.Order, error)
	CreateOrderInfo(userId uint64) responsemodels.CreateOrderInfoResponse
	GetAllHistory(userId uint) []dbmodels.Order
	ExecuteMarketOrders()
	ExecuteLimitOrderBuyOrders()
	ExecuteLimitOrderSellOrders()
	ExecuteStopOrderSellOrders()
	GetOrderList(userId uint) []responsemodels.OrderResponse
	CancelOrder(userId uint, orderId uint) error
}
