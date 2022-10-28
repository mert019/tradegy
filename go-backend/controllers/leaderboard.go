package controllers

import (
	"go-backend/controllers/response"
	"go-backend/interfaces/adaptors/httpapi"
	"go-backend/interfaces/core"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LeaderboardController struct {
	leaderboardManager core.ILeaderbaordManager
}

func NewLeaderboardController(leaderboardManager core.ILeaderbaordManager) httpapi.ILeaderboardController {

	lc := &LeaderboardController{
		leaderboardManager: leaderboardManager,
	}
	log.Printf("LeaderboardController created successfully")
	return lc
}

func (lc *LeaderboardController) GetLeaderboardList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		respList, err := lc.leaderboardManager.GetLeaderboardList()

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Error on leaderboard service", nil)
			return
		} else {
			response.JSON(w, http.StatusOK, "", respList)
			return
		}
	}
}

func (lc *LeaderboardController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1/leaderboard").Subrouter()
	subRouter.Methods(http.MethodGet).Path("/list").Handler(lc.GetLeaderboardList())
}
