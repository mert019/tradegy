package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IUserController interface {
	Register() http.HandlerFunc
	RegisterRoutes(router *mux.Router)
}
