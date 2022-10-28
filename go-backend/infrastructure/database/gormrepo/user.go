package gormrepo

import (
	"go-backend/interfaces/ports/database"
	dbmodels "go-backend/models/dbmodels"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) database.IUserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(username string, password string) (dbmodels.User, error) {
	user := dbmodels.User{UserName: username, Password: password}
	err := u.DB.Create(&user).Error
	return user, err
}

func (u *UserRepository) GetUserFromUsername(username string) (dbmodels.User, error) {
	var result dbmodels.User
	err := u.DB.Where(dbmodels.User{UserName: username}).First(&result).Error
	return result, err
}

func (u *UserRepository) GetAll() []dbmodels.User {
	var result []dbmodels.User
	u.DB.Find(&result)
	return result
}
