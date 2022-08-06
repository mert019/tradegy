package gormrepo

import (
	"go-backend/interfaces/ports/database"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"

	"gorm.io/gorm"
)

type AssetRepository struct {
	DB *gorm.DB
}

func NewAssetRepository(db *gorm.DB) database.IAssetRepository {
	return &AssetRepository{DB: db}
}

func (ar *AssetRepository) GetAll() []dbmodels.Asset {
	var result []dbmodels.Asset
	ar.DB.Find(&result)
	return result
}

func (ar *AssetRepository) GetAllCryptocurrencies() []dbmodels.Asset {
	var result []dbmodels.Asset
	ar.DB.Where(dbmodels.Asset{TypeId: enums.CRYPTOCURRENCY}).Find(&result)
	return result
}

func (ar *AssetRepository) GetByID(id uint) dbmodels.Asset {
	var result dbmodels.Asset
	ar.DB.First(&result, id)
	return result
}
