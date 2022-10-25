package requestmodels

type GetExchageRateRequest struct {
	BuyAssetID  int64 `json:"buy_asset_id"`
	SellAssetID int64 `json:"sell_asset_id"`
}

func (gerr *GetExchageRateRequest) Validate() string {
	msg := ""

	if gerr.BuyAssetID == 0 {
		msg += "Buy asset must be supplied."
	}

	if gerr.SellAssetID == 0 {
		msg += "Sell asset must be supplied."
	}

	return msg
}
