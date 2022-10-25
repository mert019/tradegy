package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IOrderController interface {
	RegisterRoutes(router *mux.Router)
	CreateOrder() http.HandlerFunc
	GetAllHistory() http.HandlerFunc
}
