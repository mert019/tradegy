package core

import (
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/requestmodels"
)

type IOrderManager interface {
	CreateOrder(order requestmodels.CreateOrderRequest, username string) (dbmodels.Order, error)
	ExecuteMarketOrderBuyOrders()
	ExecuteMarketOrderSellOrders()
	ExecuteLimitOrderBuyOrders()
	ExecuteLimitOrderSellOrders()
	ExecuteStopOrderSellOrders()
}
