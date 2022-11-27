package gormrepo

import (
	"fmt"
	"go-backend/config"
	dbmodels "go-backend/models/dbmodels"
	"log"
	"os"
	"strconv"
	"time"

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

		maxIdleConnStr := os.Getenv(config.DB_MAX_IDLE_CONNS)
		maxIdleConn, maxIdleConnErr := strconv.Atoi(maxIdleConnStr)

		maxOpenConnStr := os.Getenv(config.DB_MAX_IDLE_CONNS)
		maxOpenConn, maxOpenConnErr := strconv.Atoi(maxOpenConnStr)

		connMaxLifetimeStr := os.Getenv(config.DB_CONN_MAX_LIFETIME)
		connMaxLifetime, connMaxLifetimeErr := strconv.Atoi(connMaxLifetimeStr)

		if maxIdleConnErr != nil || maxOpenConnErr != nil || connMaxLifetimeErr != nil {
			log.Fatalln("error on getting or parsing DB_MAX_IDLE_CONNS, DB_MAX_OPEN_CONNS, DB_CONN_MAX_LIFETIME")
		}

		sqlDB, sqldbErr := db.DB()
		if sqldbErr != nil {
			log.Fatalln("sql db error", sqldbErr)
		}
		sqlDB.SetMaxIdleConns(maxIdleConn)
		sqlDB.SetMaxOpenConns(maxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
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
