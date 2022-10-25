package core

import "go-backend/models/responsemodels"

type IAssetManager interface {
	GetWealthInformationByUserId(userID uint64) ([]responsemodels.WealthInformationResponse, error)
	GetExchangeRateByAssetId(buyAssetID int64, sellAssetID int64) (float64, error)
}
