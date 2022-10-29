package tasks

import (
	"context"
	"encoding/json"
	cacheKeys "go-backend/infrastructure/cache"
	core "go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	"go-backend/models/responsemodels"
	"log"
	"math"
	"sort"
	"time"
)

type LeaderboardTask struct {
	cache        cache.ICache
	assetManager core.IAssetManager
	userManager  core.IUserManager
	ticker       *time.Ticker
	done         chan bool
}

func (lbt *LeaderboardTask) Start() {
	lbt.execute()
	log.Println("LeaderboardTask execution completed successfully")
	for {
		select {
		case <-lbt.done:
			lbt.ticker.Stop()
			return
		case <-lbt.ticker.C:
			lbt.execute()
			log.Println("LeaderboardTask execution completed successfully")
		}
	}
}

func (lbt *LeaderboardTask) execute() {

	ctx := context.Background()

	var leaderboardItems []responsemodels.LeaderboardItem
	users := lbt.userManager.GetAll()

	// Iterate users
	for _, user := range users {

		wealthInfo, err := lbt.assetManager.GetWealthInformationByUserId(uint64(user.ID))
		if err != nil {
			log.Println("LeaderboardTask: error on getting wealth info for user id: ", user.ID, ": ", err)
			continue
		}

		totalUsdWealth := 0.0
		for _, wealth := range wealthInfo {
			totalUsdWealth += wealth.UsdAmount
		}

		item := responsemodels.LeaderboardItem{
			UserName:       user.UserName,
			TotalUsdAmount: totalUsdWealth,
		}

		leaderboardItems = append(leaderboardItems, item)
		sort.Slice(leaderboardItems, func(i, j int) bool {
			return leaderboardItems[i].TotalUsdAmount > leaderboardItems[j].TotalUsdAmount
		})
		leaderboardItems = leaderboardItems[:int(math.Min(float64(len(leaderboardItems)), 10))]
	}

	// Save to cache
	leaderboardResp := responsemodels.LeaderboardResponse{
		UpdatedAt:        time.Now(),
		LeaderBoardItems: leaderboardItems,
	}
	cacheVal, jsonMarshallErr := json.Marshal(leaderboardResp)
	if jsonMarshallErr != nil {
		log.Printf("LeaderboardTask: Error on json.Marshal: %v\n", jsonMarshallErr)
	}

	lbt.cache.Set(ctx, cacheKeys.LEADERBOARD_LIST, cacheVal, 24*60*60)
}
