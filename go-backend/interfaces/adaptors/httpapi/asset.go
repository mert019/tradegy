package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IAssetController interface {
	RegisterRoutes(router *mux.Router)
	GetWealthInformation() http.HandlerFunc
	GetExchageRate() http.HandlerFunc
}
