package gormrepo

import (
	"go-backend/interfaces/ports/database"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"

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
