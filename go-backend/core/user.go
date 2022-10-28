package core

import (
	"errors"
	"go-backend/config"
	"go-backend/interfaces/core"
	"go-backend/interfaces/ports/database"
	customerrors "go-backend/models/customerrors"
	dbmodels "go-backend/models/dbmodels"
	"go-backend/models/enums"
	"go-backend/utils"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserManager struct {
	userRepository  database.IUserRepository
	orderRepository database.IOrderRepository
}

func NewUserManager(userRepository database.IUserRepository, orderRepository database.IOrderRepository) core.IUserManager {
	um := &UserManager{
		userRepository:  userRepository,
		orderRepository: orderRepository,
	}
	log.Println("UserManager created successfully")
	return um
}

func (um *UserManager) CreateUser(username string, password string) (dbmodels.User, error) {

	// Validate.
	if !utils.ValidateStringLength(username, 8, 16) {
		return dbmodels.User{}, customerrors.ErrUsernameLength
	}
	if !utils.ValidateStringLength(password, 8, 16) {
		return dbmodels.User{}, customerrors.ErrPasswordLength
	}

	// Check if user exists.
	if _, err := um.userRepository.GetUserFromUsername(username); err == nil {
		return dbmodels.User{}, customerrors.ErrUsernameAlreadyExist
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dbmodels.User{}, err
	}

	//Create user.
	hashedPassword, err := utils.GetPasswordHash(password)
	if err != nil {
		return dbmodels.User{}, err
	}
	user, userCreateErr := um.userRepository.CreateUser(username, string(hashedPassword))
	if userCreateErr != nil {
		return dbmodels.User{}, userCreateErr
	}

	initAmount, initAmountErr := strconv.ParseFloat(os.Getenv(config.USER_INIT_USD_AMOUNT), 64)
	if initAmountErr != nil {
		return dbmodels.User{}, err
	}

	// Init user money.
	order := dbmodels.Order{
		BuyAmount:         initAmount,
		SellAmount:        0,
		UserID:            user.ID,
		OrderTypeID:       enums.INITIALIZE_USER,
		BuyAssetID:        enums.USD,
		SellAssetID:       enums.USD,
		OrderStatusID:     enums.EXECUTED,
		ExecutionDateTime: time.Now(),
	}

	_, orderErr := um.orderRepository.CreateOrder(order)
	if orderErr != nil {
		return user, orderErr
	}

	return user, nil
}

func (um *UserManager) GetAll() []dbmodels.User {
	return um.userRepository.GetAll()
}
