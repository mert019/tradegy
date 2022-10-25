package database

import (
	dbmodels "go-backend/models/dbmodels"
	responsemodel "go-backend/models/responsemodels"
)

type IAssetRepository interface {
	GetAll() []dbmodels.Asset
	GetAllCryptocurrencies() []dbmodels.Asset
	GetByID(id uint) dbmodels.Asset
	GetWealthInformationByUserId(userId uint) []responsemodel.WealthInformationResponse
}
