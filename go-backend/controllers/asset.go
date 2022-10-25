package controllers

import (
	"go-backend/auth"
	"go-backend/controllers/response"
	"go-backend/interfaces/adaptors/httpapi"
	"go-backend/interfaces/core"
	"go-backend/models/requestmodels"
	"go-backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AssetController struct {
	assetManager   core.IAssetManager
	authMiddleware auth.JWTAuth
}

func NewAssetController(assetManager core.IAssetManager, authMiddleware auth.JWTAuth) httpapi.IAssetController {
	ac := &AssetController{
		assetManager:   assetManager,
		authMiddleware: authMiddleware,
	}
	log.Println("AssetController created successfuly")
	return ac
}

func (ac *AssetController) GetWealthInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user_id.
		userIDstr := r.Header.Get("user_id")
		userID, parseErr := strconv.ParseUint(userIDstr, 10, 32)
		if parseErr != nil {
			response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
			return
		}

		wealthInfo, err := ac.assetManager.GetWealthInformationByUserId(userID)

		// Response
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
		} else {
			response.JSON(w, http.StatusOK, "", wealthInfo)
		}
	}
}

func (ac *AssetController) GetExchageRate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get json
		var model requestmodels.GetExchageRateRequest
		if utils.DecodeJSONFromBody(r, &model) != nil {
			response.JSON(w, http.StatusBadRequest, "Error on decoding message.", nil)
			return
		}

		// Validate model.
		if validationMessage := model.Validate(); len(validationMessage) > 0 {
			response.JSON(w, http.StatusBadRequest, validationMessage, nil)
			return
		}

		// Get exchange rate.
		exchangeRate, err := ac.assetManager.GetExchangeRateByAssetId(model.BuyAssetID, model.SellAssetID)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error(), nil)
		} else {
			response.JSON(w, http.StatusOK, "", exchangeRate)
		}
	}
}

func (ac *AssetController) RegisterRoutes(router *mux.Router) {
	prefix := "/api/v1/asset"
	router.Methods(http.MethodPost).Path(prefix + "/exchangerate").Handler(ac.GetExchageRate())
	subRouter := router.PathPrefix(prefix).Subrouter()
	subRouter.Use(ac.authMiddleware.AuthMiddleware)
	subRouter.Methods(http.MethodGet).Path("/wealthinfo").Handler(ac.GetWealthInformation())
}
