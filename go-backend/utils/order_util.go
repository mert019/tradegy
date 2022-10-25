package utils

import (
	"context"
	"encoding/json"
	cacheKeys "go-backend/infrastructure/cache"
	"go-backend/interfaces/ports/cache"
	"go-backend/models/coingecko"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"
	"log"
	"math"
)

func GetTotalAmountFromOrdersByAssetID(orders []dbmodels.Order, assetID uint) float64 {

	buyTotal := 0.0
	sellTotal := 0.0

	for _, order := range orders {
		if order.BuyAssetID == assetID {
			buyTotal += order.BuyAmount
		}
		if order.SellAssetID == assetID {
			sellTotal += order.SellAmount
		}
	}
	return buyTotal - sellTotal
}

func RoundAmount(amount float64, digits uint) float64 {
	ratio := math.Pow(10, float64(digits))
	return math.Round(amount*ratio) / ratio
}

func GetConversionRate(cache cache.ICache, buyAsset dbmodels.Asset, sellAsset dbmodels.Asset) (float64, error) {

	ctx := context.Background()
	var buyAssetData coingecko.CryptocurrencyLatestInfoResponse
	var sellAssetData coingecko.CryptocurrencyLatestInfoResponse

	if buyAsset.ID == enums.USD && sellAsset.ID == enums.USD {
		return 1.0, nil
	}

	if buyAsset.ID != enums.USD {
		buyAssetInfo, buyInfoErr := cache.Get(ctx, cacheKeys.LATEST_CRYPTOCURRENCY_INFO_PREFIX+buyAsset.ApiId)
		if buyInfoErr != nil {
			log.Printf("Error on getting buy asset info: %v\n", buyInfoErr)
			return 0, buyInfoErr
		}
		buyAssetData = coingecko.CryptocurrencyLatestInfoResponse{}
		if err := json.Unmarshal([]byte(buyAssetInfo), &buyAssetData); err != nil {
			log.Printf("Error decoding buy asset info: %v\n", err)
			return 0, err
		}
	}
	if sellAsset.ID != enums.USD {
		sellAssetInfo, sellInfoErr := cache.Get(ctx, cacheKeys.LATEST_CRYPTOCURRENCY_INFO_PREFIX+sellAsset.ApiId)
		if sellInfoErr != nil {
			log.Printf("Error on getting sell asset info: %v\n", sellInfoErr)
			return 0, sellInfoErr
		}
		sellAssetData = coingecko.CryptocurrencyLatestInfoResponse{}
		if err := json.Unmarshal([]byte(sellAssetInfo), &sellAssetData); err != nil {
			log.Printf("Error decoding sell asset info: %v\n", err)
			return 0, err
		}
	}

	if buyAsset.ID == enums.USD {
		return sellAssetData.Current_price, nil
	} else if sellAsset.ID == enums.USD {
		return 1 / buyAssetData.Current_price, nil
	} else {
		return sellAssetData.Current_price / buyAssetData.Current_price, nil
	}
}
