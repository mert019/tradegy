package main

import (
	"flag"
	"go-backend/auth"
	"go-backend/config"
	"go-backend/controllers"
	"go-backend/core"
	"go-backend/infrastructure/cache"
	"go-backend/infrastructure/database"
	"go-backend/interfaces/adaptors/httpapi"
	coreInterfaces "go-backend/interfaces/core"
	cacheInterface "go-backend/interfaces/ports/cache"
	databaseInterface "go-backend/interfaces/ports/database"
	"go-backend/tasks"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

// Cache
var cacheObj cacheInterface.ICache

// Controllers
var userController httpapi.IUserController
var orderController httpapi.IOrderController

var jwtauth *auth.JWTAuth

// Managers
var userManager coreInterfaces.IUserManager
var orderManager coreInterfaces.IOrderManager

// Repositories
var userRepository databaseInterface.IUserRepository
var orderRepository databaseInterface.IOrderRepository
var assetRepository databaseInterface.IAssetRepository

func init() {

	// Flags
	envFile := flag.String("env", "./config/.env", "")
	flag.Parse()

	// Initialize environment variables
	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatalf("Error on loading .env file %s\n", err.Error())
	}

	// Initialize Router
	router = mux.NewRouter()

	// Initialize Repositories
	database.InitRepositoryPackage()
	userRepository = database.GetUserRepository()
	orderRepository = database.GetOrderRepository()
	assetRepository = database.GetAssetRepository()

	// Initialize Cache
	cache.InitializeCache()
	cacheObj = cache.GetCache()

	// Initialize Managers
	userManager = core.NewUserManager(userRepository, orderRepository)
	orderManager = core.NewOrderManager(userRepository, orderRepository, assetRepository, cacheObj)

	// Initialize Controllers
	jwtauth = auth.NewJWTAuth(userRepository)
	userController = controllers.NewUserController(userManager)
	orderController = controllers.NewOrderController(orderManager, *jwtauth)

	// Register Routes
	jwtauth.AddRoute(router)
	userController.RegisterRoutes(router)
	orderController.RegisterRoutes(router)

	// Initialize and Start Tasks
	if isTasksEnabled, err := strconv.ParseBool(os.Getenv(config.TASKS_ENABLED)); err != nil {
		log.Fatalf("Error on TASKS_ENABLED environment variable: %v\n", err)
	} else if isTasksEnabled {
		tasks.InitializeTasks(cacheObj, assetRepository, orderManager)
		tasks.Start()
	}
}

func main() {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	log.Printf("Starting on port %s\n", os.Getenv(config.SERVER_PORT))
	log.Fatalln(http.ListenAndServe(":"+os.Getenv(config.SERVER_PORT), handlers.CORS(headers, methods, origins, credentials)(router)))
}
