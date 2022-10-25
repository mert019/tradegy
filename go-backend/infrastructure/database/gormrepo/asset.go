package gormrepo

import (
	"go-backend/interfaces/ports/database"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"
	responsemodel "go-backend/models/responsemodels"

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

func (ar *AssetRepository) GetWealthInformationByUserId(userId uint) []responsemodel.WealthInformationResponse {

	query := `SELECT buy_asset_id asset_id, COALESCE(buy_sum, 0) - COALESCE(sell_sum, 0) Amount, assets.name Name, assets.image_source
			FROM 
				(
					SELECT o.buy_asset_id, SUM(o.buy_amount) buy_sum FROM users u 
					INNER JOIN orders o ON u.id = o.user_id 
					WHERE 
						u.id = ?
						AND o.order_status_id = ? 
						AND u.deleted_at IS NULL AND o.deleted_at IS NULL 
						GROUP BY o.buy_asset_id) 
						as buy_sum 
				FULL JOIN 
						(SELECT o.sell_asset_id, SUM(o.sell_amount) sell_sum 
						FROM users u INNER JOIN orders o ON u.id = o.user_id 
						WHERE u.id = ? AND o.order_status_id = ? AND u.deleted_at IS NULL AND o.deleted_at IS NULL 
						GROUP BY o.sell_asset_id) as sell_sum 
						ON buy_sum.buy_asset_id = sell_sum.sell_asset_id 
				LEFT JOIN assets assets ON buy_sum.buy_asset_id = assets.id`

	var retVal []responsemodel.WealthInformationResponse
	ar.DB.Raw(query, userId, enums.EXECUTED, userId, enums.EXECUTED).Scan(&retVal)
	return retVal
}
