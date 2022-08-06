package database

import dbmodels "go-backend/models/dbmodels"

type IAssetRepository interface {
	GetAll() []dbmodels.Asset
	GetAllCryptocurrencies() []dbmodels.Asset
	GetByID(id uint) dbmodels.Asset
}
