package core

import dbmodels "go-backend/models/dbmodels"

type IUserManager interface {
	CreateUser(username string, password string) (dbmodels.User, error)
	GetAll() []dbmodels.User
}
