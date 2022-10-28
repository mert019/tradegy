package responsemodels

import "time"

type OrderResponse struct {
	OrderId              uint      `json:"order_id"`
	BuyAmount            float64   `json:"buy_amount"`
	SellAmount           float64   `json:"sell_amount"`
	BuyAssetCode         string    `json:"buy_asset_code"`
	SellAssetCode        string    `json:"sell_asset_code"`
	BuyAssetImageSource  string    `json:"buy_asset_image_source"`
	SellAssetImageSource string    `json:"sell_asset_image_source"`
	Limit                float64   `json:"limit"`
	OrderType            string    `json:"order_type"`
	OrderStatus          string    `json:"order_status"`
	CreatedAt            time.Time `json:"created_at"`
	ExecutedAt           time.Time `json:"executed_at"`
}
