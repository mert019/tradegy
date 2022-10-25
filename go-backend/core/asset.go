package core

import (
	"go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	"go-backend/interfaces/ports/database"
	customerrors "go-backend/models/customerrors"
	"go-backend/models/enums"
	"go-backend/models/responsemodels"
	"go-backend/utils"
	"log"
)

type AssetManager struct {
	orderRepository database.IOrderRepository
	assetRepository database.IAssetRepository
	cache           cache.ICache
}

func NewAssetManager(orderRepository database.IOrderRepository, cache cache.ICache, assetRepository database.IAssetRepository) core.IAssetManager {
	am := &AssetManager{
		orderRepository: orderRepository,
		cache:           cache,
		assetRepository: assetRepository,
	}
	log.Println("AssetManager created successfully")
	return am
}

func (am *AssetManager) GetWealthInformationByUserId(userID uint64) ([]responsemodels.WealthInformationResponse, error) {

	infoData := am.assetRepository.GetWealthInformationByUserId(uint(userID))

	usdAsset := am.assetRepository.GetByID(enums.USD)

	for index, info := range infoData {
		asset := am.assetRepository.GetByID(uint(info.AssetId))
		if rate, err := utils.GetConversionRate(am.cache, usdAsset, asset); err == nil {
			infoData[index].UsdAmount = rate * info.Amount
		} else {
			return nil, err
		}
	}

	return infoData, nil
}

func (am *AssetManager) GetExchangeRateByAssetId(buyAssetID int64, sellAssetID int64) (float64, error) {

	buyAsset := am.assetRepository.GetByID(uint(buyAssetID))
	sellAsset := am.assetRepository.GetByID(uint(sellAssetID))

	if buyAsset.ID == 0 || sellAsset.ID == 0 {
		return 0, customerrors.ErrInvalidAssetID
	}

	if rate, err := utils.GetConversionRate(am.cache, buyAsset, sellAsset); err == nil {
		return rate, nil
	} else {
		return 0, err
	}
}
