package controllers

import (
	"go-backend/auth"
	"go-backend/controllers/response"
	"go-backend/interfaces/core"
	"go-backend/models/requestmodels"
	"go-backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderController struct {
	orderManager   core.IOrderManager
	authMiddleware auth.JWTAuth
}

func NewOrderController(orderManager core.IOrderManager, authMiddleware auth.JWTAuth) *OrderController {
	oc := &OrderController{orderManager: orderManager}
	log.Println("OrderController created successfuly")
	return oc
}

func (oc *OrderController) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get username.
		username := r.Header.Get("username")

		// Get json.
		var model requestmodels.CreateOrderRequest
		if utils.DecodeJSONFromBody(r, &model) != nil {
			response.JSON(w, http.StatusBadRequest, "Error on decoding message.", nil)
			return
		}

		// Validate model.
		if validationMessage := model.Validate(); len(validationMessage) > 0 {
			response.JSON(w, http.StatusBadRequest, validationMessage, nil)
			return
		}

		// Create Order.
		_, newOrderErr := oc.orderManager.CreateOrder(model, username)

		// Response.
		if newOrderErr != nil {
			response.JSON(w, http.StatusInternalServerError, newOrderErr.Error(), nil)
			return
		} else {
			response.JSON(w, http.StatusCreated, "Order created successfuly.", nil)
			return
		}
	}
}

func (oc *OrderController) CreateOrderInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get user_id
		userIDstr := r.Header.Get("user_id")
		userId, parseErr := strconv.ParseUint(userIDstr, 10, 32)
		if parseErr != nil {
			response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
			return
		}

		// Get Info
		info := oc.orderManager.CreateOrderInfo(userId)

		// Response.
		response.JSON(w, http.StatusOK, "", info)
	}
}

func (oc *OrderController) GetAllHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get user_id
		userIDstr := r.Header.Get("user_id")
		userId, parseErr := strconv.ParseUint(userIDstr, 10, 32)
		if parseErr != nil {
			response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
			return
		}

		// Get history
		history := oc.orderManager.GetAllHistory(uint(userId))
		response.JSON(w, http.StatusOK, "", history)

	}
}

func (oc *OrderController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1/order").Subrouter()
	subRouter.Use(oc.authMiddleware.AuthMiddleware)
	subRouter.Methods(http.MethodPost).Path("/create").Handler(oc.CreateOrder())
	subRouter.Methods(http.MethodGet).Path("/createinfo").Handler(oc.CreateOrderInfo())
	subRouter.Methods(http.MethodGet).Path("/allhistory").Handler(oc.GetAllHistory())
}
