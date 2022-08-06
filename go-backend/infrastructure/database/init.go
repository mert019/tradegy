package database

import (
	"go-backend/config"
	"go-backend/infrastructure/database/gormrepo"
	"go-backend/interfaces/ports/database"
	"log"
	"os"
)

// Constants
const (
	GORMREPO string = "gormrepo"
)

var repositoryType string

// Repositories
var userRepository database.IUserRepository
var orderRepository database.IOrderRepository
var assetRepository database.IAssetRepository

func InitRepositoryPackage() {
	repositoryType = os.Getenv(config.REPOSITORY_TYPE)
	if repositoryType == GORMREPO {
		gormrepo.GetDatabase()
		log.Printf("%s repository initialized successfully\n", repositoryType)
	} else {
		log.Fatalln("REPOSITORY_TYPE could not match.")
	}

}

// Getters
func GetUserRepository() database.IUserRepository {
	if userRepository == nil {
		if repositoryType == GORMREPO {
			db := gormrepo.GetDatabase()
			userRepository = gormrepo.NewUserRepository(db)
			log.Println("userRepository initialized successfully")
		}
	}
	return userRepository
}

func GetOrderRepository() database.IOrderRepository {
	if orderRepository == nil {
		if repositoryType == GORMREPO {
			db := gormrepo.GetDatabase()
			orderRepository = gormrepo.NewOrderRepository(db)
			log.Println("orderRepository initialized successfully")
		}
	}
	return orderRepository
}

func GetAssetRepository() database.IAssetRepository {
	if assetRepository == nil {
		if repositoryType == GORMREPO {
			db := gormrepo.GetDatabase()
			assetRepository = gormrepo.NewAssetRepository(db)
			log.Println("assetRepository initialized successfully")
		}
	}
	return assetRepository
}
