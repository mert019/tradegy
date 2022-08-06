package tasks

import (
	"context"
	"encoding/json"
	cacheKeys "go-backend/infrastructure/cache"
	"go-backend/interfaces/ports/cache"
	databaseInterface "go-backend/interfaces/ports/database"
	"go-backend/models/coingecko"
	models "go-backend/models/dbmodels"
	"go-backend/models/enums"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CryptocurrencyInfoTask struct {
	cache           cache.ICache
	assetRepository databaseInterface.IAssetRepository
	ticker          *time.Ticker
	done            chan bool
}

func (cit *CryptocurrencyInfoTask) Start() {
	cit.execute()
	for {
		select {
		case <-cit.done:
			cit.ticker.Stop()
			return
		case <-cit.ticker.C:
			cit.execute()
			log.Println("CryptocurrencyInfoTask execution completed successfully")
		}
	}
}

func (cit *CryptocurrencyInfoTask) execute() {

	ctx := context.Background()

	allCC := cit.assetRepository.GetAllCryptocurrencies()
	if len(allCC) == 0 {
		log.Println("assetRepository.GetAllCryptocurrencies returned an empty list.")
		return
	}

	getUSD := cit.assetRepository.GetByID(enums.USD)
	if getUSD.ID == 0 {
		log.Println("assetRepository returned an empty value for USD asset")
		return
	}

	url := cit.buildUrl(allCC, getUSD)

	client := http.Client{
		Timeout: time.Second * 25,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Printf("Error creating response: %v\n", getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Printf("Error on reading body: %v\n", readErr)
	}

	var info []coingecko.CryptocurrencyLatestInfoResponse
	jsonErr := json.Unmarshal(body, &info)
	if jsonErr != nil {
		log.Printf("Error on json.Unmarshal: %v\n", jsonErr)
	}

	for _, element := range info {
		val, jsonMarshallErr := json.Marshal(element)
		if jsonMarshallErr != nil {
			log.Printf("Error on json.Marshal: %v\n", jsonMarshallErr)
			continue
		}
		setErr := cit.cache.Set(ctx, cacheKeys.LATEST_CRYPTOCURRENCY_INFO_PREFIX+element.ID, val, 80)
		if setErr != nil {
			log.Printf("Error on setting cache: %v\n", setErr)
		}
	}

}

func (cit *CryptocurrencyInfoTask) buildUrl(ccs []models.Asset, usd models.Asset) string {
	url := "https://api.coingecko.com/api/v3/coins/markets?per_page=250"

	// Add crypto currency ids.
	url += "&ids="
	for _, elem := range ccs {
		url += elem.ApiId + "%2C"
	}

	// Add vs_currency
	url += "&vs_currency=" + usd.ApiId

	return url
}
