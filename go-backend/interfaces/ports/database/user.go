package database

import dbmodels "go-backend/models/dbmodels"

type IUserRepository interface {
	CreateUser(username string, password string) (dbmodels.User, error)
	GetUserFromUsername(username string) (dbmodels.User, error)
}
