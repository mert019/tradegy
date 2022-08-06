package gormrepo

import (
	"fmt"
	"go-backend/config"
	dbmodels "go-backend/models/dbmodels"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {

	if db == nil {

		server := os.Getenv(config.DB_SERVER)
		port := os.Getenv(config.DB_PORT)
		user := os.Getenv(config.DB_USER)
		password := os.Getenv(config.DB_PASSWORD)
		database := os.Getenv(config.DB_NAME)

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			server, user, password, database, port)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Error creating database: %v", err)
		} else {
			log.Println("Database created successfully")
		}
	}

	return db
}

func MigrateTables() {
	db := GetDatabase()
	err := db.AutoMigrate(
		&dbmodels.Enum{},
		&dbmodels.Asset{},
		&dbmodels.User{},
		&dbmodels.Order{},
	)
	if err != nil {
		log.Fatalf("Error on migrating tables: %v", err)
	}
}
