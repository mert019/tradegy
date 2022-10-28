package core

import "go-backend/models/responsemodels"

type ILeaderbaordManager interface {
	GetLeaderboardList() (responsemodels.LeaderboardResponse, error)
}
