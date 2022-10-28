package core

import (
	"context"
	"encoding/json"
	cacheKeys "go-backend/infrastructure/cache"
	"go-backend/interfaces/core"
	"go-backend/interfaces/ports/cache"
	"go-backend/models/responsemodels"
	"log"
)

type LeaderboardManager struct {
	cache cache.ICache
}

func NewLeaderboardManager(cache cache.ICache) core.ILeaderbaordManager {
	lm := &LeaderboardManager{
		cache: cache,
	}
	log.Println("LeaderboardManager created successfully")
	return lm
}

func (lm *LeaderboardManager) GetLeaderboardList() (responsemodels.LeaderboardResponse, error) {

	ctx := context.Background()

	var retVal responsemodels.LeaderboardResponse

	leaderboardData, cacheGetErr := lm.cache.Get(ctx, cacheKeys.LEADERBOARD_LIST)
	if cacheGetErr != nil {
		log.Println("Error on getting leaderboard list from cache", cacheGetErr)
		return retVal, cacheGetErr
	}

	if err := json.Unmarshal([]byte(leaderboardData), &retVal); err != nil {
		log.Printf("Error decoding leaderboard list: %v\n", err)
		return retVal, err
	}

	return retVal, nil
}
