package controllers

import (
	"errors"
	"go-backend/controllers/response"
	"go-backend/interfaces/adaptors/httpapi"
	"go-backend/interfaces/core"
	exceptions "go-backend/models/customerrors"
	"go-backend/models/requestmodels"
	"go-backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	userManager core.IUserManager
}

func NewUserController(userManager core.IUserManager) httpapi.IUserController {
	uc := &UserController{userManager: userManager}
	log.Println("UserController created successfuly")
	return uc
}

func (us *UserController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get json
		var model requestmodels.UserRegisterRequest
		if utils.DecodeJSONFromBody(r, &model) != nil {
			response.JSON(w, http.StatusBadRequest, "Error on decoding message.", nil)
			return
		}

		// Validate model.
		if validationMessage := model.Validate(); len(validationMessage) > 0 {
			response.JSON(w, http.StatusBadRequest, validationMessage, nil)
			return
		}

		// Create user.
		username := model.UserName
		password := model.Password

		_, err := us.userManager.CreateUser(username, password)

		// Response
		if err != nil && errors.Is(err, exceptions.ErrUsernameAlreadyExist) {
			response.JSON(w, http.StatusConflict, err.Error(), nil)
			return
		} else if err != nil && (errors.Is(err, exceptions.ErrUsernameLength) || errors.Is(err, exceptions.ErrPasswordLength)) {
			response.JSON(w, http.StatusBadRequest, err.Error(), nil)
			return
		} else if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Oops! Something went wrong.", nil)
			return
		} else {
			response.JSON(w, http.StatusCreated, "User created successfuly.", nil)
			return
		}
	}
}

func (uc *UserController) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1/user").Subrouter()
	subRouter.Methods(http.MethodPost).Path("/register").Handler(uc.Register())
}
