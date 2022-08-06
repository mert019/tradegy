package auth

import (
	"errors"
	"go-backend/controllers/response"
	"go-backend/interfaces/ports/database"
	requestmodels "go-backend/models/requestmodels"
	"go-backend/models/responsemodels"
	utils "go-backend/utils"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type JWTAuth struct {
	UserRepository database.IUserRepository
}

func NewJWTAuth(userRepository database.IUserRepository) *JWTAuth {
	ja := &JWTAuth{UserRepository: userRepository}
	log.Println("JWTAuth created successfully")
	return ja
}

func (jwtauth *JWTAuth) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get json from request.
		var model requestmodels.UserLoginRequest
		if utils.DecodeJSONFromBody(r, &model) != nil {
			response.JSON(w, http.StatusBadRequest, "Error on decoding message.", nil)
			return
		}

		// Validate model.
		if validationMessage := model.Validate(); len(validationMessage) > 0 {
			response.JSON(w, http.StatusBadRequest, validationMessage, nil)
			return
		}

		// Get user from database.
		username := model.UserName
		password := model.Password
		user, err := jwtauth.UserRepository.GetUserFromUsername(username)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			response.JSON(w, http.StatusNotFound, "Invalid login credentials.", nil)
			return
		} else if err != nil {
			response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
			return
		}

		// Check password.
		expectedPassword := []byte(user.Password)
		if bcrypt.CompareHashAndPassword(expectedPassword, []byte(password)) != nil {
			response.JSON(w, http.StatusUnauthorized, "Invalid login credentials.", nil)
			return
		} else {
			// Create token.
			token, err := utils.GetToken(user.UserName)
			if err != nil {
				response.JSON(w, http.StatusInternalServerError, "Ooops, something went wrong.", nil)
			} else {
				w.Header().Set("Authorization", "Bearer "+token)
				responseModel := responsemodels.UserLoginResponse{Token: token}
				response.JSON(w, http.StatusOK, "Login Successful.", responseModel)
				return
			}
		}
	}
}

// Login required auth middleware.
func (jwtauth *JWTAuth) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			response.JSON(w, http.StatusUnauthorized, "Login token not found. Login required.", nil)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "Error on verifying user.", nil)
			return
		}
		//Pass jwt parameters.
		username := claims.(jwt.MapClaims)["username"].(string)
		r.Header.Set("username", username)

		next.ServeHTTP(w, r)
	})
}

func (jwtauth *JWTAuth) AddRoute(router *mux.Router) {
	subRouter := router.PathPrefix("/api/v1/auth").Subrouter()
	subRouter.Methods(http.MethodPost).Path("/login").Handler(jwtauth.Authenticate()).Name("login")
}
