package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ILeaderboardController interface {
	RegisterRoutes(router *mux.Router)
	GetLeaderboardList() http.HandlerFunc
}
