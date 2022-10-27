package core

import (
	"go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	"go-backend/interfaces/ports/database"
	customerrors "go-backend/models/customerrors"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"
	"go-backend/models/requestmodels"
	"go-backend/models/responsemodels"
	"go-backend/utils"
	"log"
	"time"
)

type OrderManager struct {
	userRepository  database.IUserRepository
	orderRepository database.IOrderRepository
	assetRepository database.IAssetRepository
	cache           cache.ICache
}

func NewOrderManager(userRepository database.IUserRepository, orderRepository database.IOrderRepository, assetRepository database.IAssetRepository, cache cache.ICache) core.IOrderManager {
	om := &OrderManager{
		userRepository:  userRepository,
		orderRepository: orderRepository,
		assetRepository: assetRepository,
		cache:           cache,
	}
	log.Println("OrderManager created successfully")
	return om
}

func (om *OrderManager) CreateOrder(orderReq requestmodels.CreateOrderRequest, username string) (dbmodels.Order, error) {

	// Get User.
	user, err := om.userRepository.GetUserFromUsername(username)
	if err != nil {
		log.Println("Error on getting user: ", err)
		return dbmodels.Order{}, err
	}

	order := dbmodels.Order{
		BuyAssetID:    orderReq.BuyAssetID,
		SellAssetID:   orderReq.SellAssetID,
		OrderTypeID:   orderReq.OrderTypeID,
		OrderStatusID: enums.OPEN,
		UserID:        user.ID,
		Limit:         orderReq.Limit,
	}

	// Validate BuyAssetID and SellAssetID.
	buyAsset := om.assetRepository.GetByID(orderReq.BuyAssetID)
	sellAsset := om.assetRepository.GetByID(orderReq.SellAssetID)
	if buyAsset.ID == 0 || sellAsset.ID == 0 {
		return dbmodels.Order{}, customerrors.ErrInvalidAssetID
	}

	// Validate Amount.
	sellAssetOrders := om.orderRepository.GetAmountEffectiveOrdersByAssetID(orderReq.SellAssetID, user.ID)

	totalSellAsset := utils.GetTotalAmountFromOrdersByAssetID(sellAssetOrders, orderReq.SellAssetID)

	if totalSellAsset-orderReq.Amount < 0 {
		return dbmodels.Order{}, customerrors.ErrInsufficientAssetBalance
	}

	order.SellAmount = orderReq.Amount

	return om.orderRepository.CreateOrder(order)
}

func (om *OrderManager) CreateOrderInfo(userId uint64) responsemodels.CreateOrderInfoResponse {

	retVal := responsemodels.CreateOrderInfoResponse{}

	assets := om.assetRepository.GetAll()
	retVal.BuyAssets = assets

	sellAssets := om.orderRepository.GetSellableAssetsByUserId(userId)
	retVal.SellAssets = sellAssets

	return retVal
}

func (om *OrderManager) GetAllHistory(userId uint) []dbmodels.Order {

	return om.orderRepository.GetOrdersByUserId(userId)
}

func (om *OrderManager) GetOrderList(userId uint) []responsemodels.OrderResponse {

	return om.orderRepository.GetOrderList(userId)
}

func (om *OrderManager) CancelOrder(userId uint, orderId uint) error {

	// Get order
	order := om.orderRepository.GetByID(orderId)
	if order.ID == 0 {
		return customerrors.ErrOrderNotFound
	}

	// Check ownership
	if order.UserID != userId {
		return customerrors.ErrUnauthorizedAccess
	}

	// Check status
	if order.OrderStatusID != enums.OPEN {
		return customerrors.ErrOrderCannotBeCancelled
	}

	// Cancel
	order.OrderStatusID = enums.CANCELLED_BY_USER

	_, updateErr := om.orderRepository.InsertUpdateOrder(order)
	return updateErr
}

func (om *OrderManager) ExecuteMarketOrderBuyOrders() {
	orders := om.orderRepository.GetOpenOrdersByOrderType(enums.MARKET_ORDER_BUY)
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		buyAsset := om.assetRepository.GetByID(order.BuyAssetID)
		sellAsset := om.assetRepository.GetByID(order.SellAssetID)
		exchangeRate, err := utils.GetConversionRate(om.cache, buyAsset, sellAsset)
		if err != nil {
			continue
		}
		amount := order.SellAmount * exchangeRate
		order.BuyAmount = utils.RoundAmount(amount, 8)
		order.ExecutionDateTime = time.Now()
		order.OrderStatusID = enums.EXECUTED
		om.orderRepository.InsertUpdateOrder(order)
	}
}

func (om *OrderManager) ExecuteMarketOrderSellOrders() {
	orders := om.orderRepository.GetOpenOrdersByOrderType(enums.MARKET_ORDER_SELL)
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		buyAsset := om.assetRepository.GetByID(order.BuyAssetID)
		sellAsset := om.assetRepository.GetByID(order.SellAssetID)
		exchangeRate, err := utils.GetConversionRate(om.cache, buyAsset, sellAsset)
		if err != nil {
			continue
		}
		amount := order.SellAmount * exchangeRate
		order.BuyAmount = utils.RoundAmount(amount, 8)
		order.ExecutionDateTime = time.Now()
		order.OrderStatusID = enums.EXECUTED
		om.orderRepository.InsertUpdateOrder(order)
	}
}

func (om *OrderManager) ExecuteLimitOrderBuyOrders() {
	orders := om.orderRepository.GetOpenOrdersByOrderType(enums.LIMIT_ORDER_BUY)
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		buyAsset := om.assetRepository.GetByID(order.BuyAssetID)
		sellAsset := om.assetRepository.GetByID(order.SellAssetID)
		exchangeRate, err := utils.GetConversionRate(om.cache, buyAsset, sellAsset)
		if err != nil {
			continue
		}
		if exchangeRate >= order.Limit {
			continue
		}
		amount := order.SellAmount * exchangeRate
		order.BuyAmount = utils.RoundAmount(amount, 8)
		order.ExecutionDateTime = time.Now()
		order.OrderStatusID = enums.EXECUTED
		om.orderRepository.InsertUpdateOrder(order)
	}
}

func (om *OrderManager) ExecuteLimitOrderSellOrders() {
	orders := om.orderRepository.GetOpenOrdersByOrderType(enums.LIMIT_ORDER_SELL)
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		buyAsset := om.assetRepository.GetByID(order.BuyAssetID)
		sellAsset := om.assetRepository.GetByID(order.SellAssetID)
		exchangeRate, err := utils.GetConversionRate(om.cache, buyAsset, sellAsset)
		if err != nil {
			continue
		}
		if exchangeRate <= order.Limit {
			continue
		}
		amount := order.SellAmount * exchangeRate
		order.BuyAmount = utils.RoundAmount(amount, 8)
		order.ExecutionDateTime = time.Now()
		order.OrderStatusID = enums.EXECUTED
		om.orderRepository.InsertUpdateOrder(order)
	}
}

func (om *OrderManager) ExecuteStopOrderSellOrders() {
	orders := om.orderRepository.GetOpenOrdersByOrderType(enums.STOP_ORDER_SELL)
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		buyAsset := om.assetRepository.GetByID(order.BuyAssetID)
		sellAsset := om.assetRepository.GetByID(order.SellAssetID)
		exchangeRate, err := utils.GetConversionRate(om.cache, buyAsset, sellAsset)
		if err != nil {
			continue
		}
		if exchangeRate >= order.Limit {
			continue
		}
		amount := order.SellAmount * exchangeRate
		order.BuyAmount = utils.RoundAmount(amount, 8)
		order.ExecutionDateTime = time.Now()
		order.OrderStatusID = enums.EXECUTED
		om.orderRepository.InsertUpdateOrder(order)
	}
}
