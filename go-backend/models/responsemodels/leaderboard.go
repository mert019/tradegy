package responsemodels

import "time"

type LeaderboardResponse struct {
	UpdatedAt        time.Time         `json:"updated_at"`
	LeaderBoardItems []LeaderboardItem `json:"leaderboard_items"`
}

type LeaderboardItem struct {
	UserName       string  `json:"user_name"`
	TotalUsdAmount float64 `json:"total_usd_amount"`
}
