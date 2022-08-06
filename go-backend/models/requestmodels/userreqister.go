package requestmodels

import "go-backend/utils"

type UserRegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRegisterRequest) Validate() string {
	msg := ""
	if !utils.ValidateStringLength(u.UserName, 0, 128) {
		msg += "Username should be supplied. "
	}
	if !utils.ValidateStringLength(u.Password, 0, 128) {
		msg += "Password should be supplied. "
	}
	return msg
}
