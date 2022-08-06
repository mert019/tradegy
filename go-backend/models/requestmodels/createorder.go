package requestmodels

import "go-backend/models/enums"

type CreateOrderRequest struct {
	Amount      float64 `json:"amount"`
	Limit       float64 `json:"limit"`
	OrderTypeID int64   `json:"order_type_id"`
	BuyAssetID  uint    `json:"buy_asset_id"`
	SellAssetID uint    `json:"sell_asset_id"`
}

func (co *CreateOrderRequest) Validate() string {
	msg := ""
	if co.Amount <= 0 {
		msg += "Amount must be a positive number. "
	}
	if co.OrderTypeID < 10000 || co.OrderTypeID > 10005 {
		msg += "Invalid order type ID. "
	}
	if co.BuyAssetID <= 0 {
		msg += "Buy AssetID must be supplied. "
	}
	if co.SellAssetID <= 0 {
		msg += "Sell AssetID must be supplied. "
	}
	if co.BuyAssetID == co.SellAssetID {
		msg += "Buy asset must be different from sell asset. "
	}
	// Validate limit.
	if (co.OrderTypeID == enums.LIMIT_ORDER_BUY || co.OrderTypeID == enums.LIMIT_ORDER_SELL || co.OrderTypeID == enums.STOP_ORDER_SELL) && (co.Limit == 0) {
		msg += "Limit value does not exists. "
	}

	return msg
}
