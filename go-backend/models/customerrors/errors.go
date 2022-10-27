package exceptions

import "errors"

var (
	ErrUsernameAlreadyExist = errors.New("username already exists")

	ErrUsernameLength = errors.New("username length should be between 8 and 16 characters")

	ErrPasswordLength = errors.New("password length should be between 8 and 16 characters")

	ErrLimitDoesNotExists = errors.New("limit value does not exist")

	ErrInvalidAssetID = errors.New("invalid asset ID")

	ErrInsufficientAssetBalance = errors.New("insufficient asset balance")

	ErrOrderNotFound = errors.New("order not found")

	ErrUnauthorizedAccess = errors.New("unauthorized access")

	ErrOrderCannotBeCancelled = errors.New("order cannot be cancelled")
)
