package gormrepo

import (
	"database/sql"
	"go-backend/interfaces/ports/database"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"
	"go-backend/models/responsemodels"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) database.IOrderRepository {
	return &OrderRepository{DB: db}
}

func (o *OrderRepository) CreateOrder(order dbmodels.Order) (dbmodels.Order, error) {
	err := o.DB.Create(&order).Error
	return order, err
}

func (o *OrderRepository) GetAmountEffectiveOrdersByAssetID(asssetID uint, userID uint) []dbmodels.Order {
	var result []dbmodels.Order
	o.DB.Where(dbmodels.Order{UserID: userID}).
		Where(o.DB.Where(dbmodels.Order{BuyAssetID: asssetID}).Or(dbmodels.Order{SellAssetID: asssetID})).
		Where(o.DB.Where(dbmodels.Order{OrderStatusID: enums.OPEN}).Or(dbmodels.Order{OrderStatusID: enums.EXECUTED})).Find(&result)
	return result
}

func (o *OrderRepository) GetOpenOrdersByOrderType(orderTypeID int) []dbmodels.Order {
	var result []dbmodels.Order
	o.DB.Where(dbmodels.Order{OrderTypeID: int64(orderTypeID)}).
		Where(dbmodels.Order{OrderStatusID: enums.OPEN}).Find(&result)
	return result
}

func (o *OrderRepository) InsertUpdateOrder(order dbmodels.Order) (dbmodels.Order, error) {
	if order.ID == 0 {
		err := o.DB.Create(&order).Error
		return order, err
	} else {
		err := o.DB.Model(&order).Updates(order).Error
		return order, err
	}
}

func (o *OrderRepository) GetExecutedOrdersByUserId(userId uint) []dbmodels.Order {
	var result []dbmodels.Order
	o.DB.Where(dbmodels.Order{UserID: userId}).
		Where(dbmodels.Order{OrderStatusID: enums.EXECUTED}).Find(&result)
	return result
}

func (o *OrderRepository) GetOrdersByUserId(userId uint) []dbmodels.Order {
	var result []dbmodels.Order
	o.DB.Where(dbmodels.Order{UserID: userId}).Find(&result)
	return result
}

func (o *OrderRepository) GetSellableAssetsByUserId(userId uint64) []responsemodels.SellAssetsResponse {
	query := `
		SELECT 
			COALESCE(buy_assets.buy_amount, 0) - COALESCE(sell_assets.sell_amount, 0) available_amount, buy_asset_id asset_id, assets.code, assets.name

		FROM
			(SELECT
				orders.buy_asset_id, SUM(orders.buy_amount) buy_amount
			FROM
				public.orders orders
			WHERE
				orders.deleted_at IS NULL
				AND orders.order_status_id IN (@orderOpen, @orderExecuted)
				AND orders.user_id = @userId
			GROUP BY
				orders.buy_asset_id) buy_assets
				
			FULL JOIN 
				(SELECT
					orders.sell_asset_id, SUM(orders.sell_amount) sell_amount
				FROM
					public.orders orders
				WHERE
					orders.deleted_at IS NULL
					AND orders.order_status_id IN (@orderOpen, @orderExecuted)
					AND orders.user_id = @userId
				GROUP BY
					orders.sell_asset_id) sell_assets ON buy_assets.buy_asset_id = sell_assets.sell_asset_id
					
			INNER JOIN public.assets assets ON assets.id = buy_asset_id
		WHERE
			COALESCE(buy_assets.buy_amount, 0) - COALESCE(sell_assets.sell_amount, 0) > 0`

	var retVal []responsemodels.SellAssetsResponse
	o.DB.Raw(query, sql.Named("userId", userId), sql.Named("orderExecuted", enums.EXECUTED), sql.Named("orderOpen", enums.OPEN)).Scan(&retVal)
	return retVal
}
