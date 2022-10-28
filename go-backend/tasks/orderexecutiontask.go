package tasks

import (
	core "go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	"log"
	"time"
)

type OrderExecutionTask struct {
	cache        cache.ICache
	orderManager core.IOrderManager
	ticker       *time.Ticker
	done         chan bool
}

func (oet *OrderExecutionTask) Start() {
	for {
		select {
		case <-oet.done:
			oet.ticker.Stop()
			return
		case <-oet.ticker.C:
			oet.execute()
			log.Println("OrderExecutionTask execution completed successfully")
		}
	}
}

func (oet *OrderExecutionTask) execute() {
	oet.orderManager.ExecuteLimitOrderBuyOrders()
	oet.orderManager.ExecuteLimitOrderSellOrders()
	oet.orderManager.ExecuteMarketOrders()
	oet.orderManager.ExecuteStopOrderSellOrders()
}
