package tasks

import (
	core "go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	databaseInterface "go-backend/interfaces/ports/database"
	"time"
)

// Tasks
var done chan bool
var cryptocurrencyInfoTask *CryptocurrencyInfoTask
var orderExecuitonTask *OrderExecutionTask

func InitializeTasks(cache cache.ICache, assetRepository databaseInterface.IAssetRepository, orderManager core.IOrderManager) {

	cryptocurrencyInfoTask = &CryptocurrencyInfoTask{
		done:            done,
		ticker:          time.NewTicker(20 * time.Second),
		cache:           cache,
		assetRepository: assetRepository,
	}

	orderExecuitonTask = &OrderExecutionTask{
		done:         done,
		ticker:       time.NewTicker(20 * time.Second),
		cache:        cache,
		orderManager: orderManager,
	}
}

func Start() {
	go cryptocurrencyInfoTask.Start()
	go orderExecuitonTask.Start()
}

func Stop() {
	done <- true
}
