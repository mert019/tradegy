package responsemodels

import dbmodels "go-backend/models/dbmodels"

type CreateOrderInfoResponse struct {
	BuyAssets  []dbmodels.Asset     `json:"buy_assets"`
	SellAssets []SellAssetsResponse `json:"sell_assets"`
}

type SellAssetsResponse struct {
	AssetId         uint    `json:"asset_id"`
	Name            string  `json:"name"`
	Code            string  `json:"code"`
	AvailableAmount float64 `json:"available_amount"`
}
