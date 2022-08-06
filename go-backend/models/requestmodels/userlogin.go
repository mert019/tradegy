package requestmodels

import "go-backend/utils"

type UserLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLoginRequest) Validate() string {
	msg := ""
	if !utils.ValidateStringLength(u.UserName, 0, 16) {
		msg += "Username legnth must be between 0 and 16 characters. "
	}
	if !utils.ValidateStringLength(u.Password, 0, 16) {
		msg += "Password legnth must be between 0 and 16 characters. "
	}
	return msg
}
