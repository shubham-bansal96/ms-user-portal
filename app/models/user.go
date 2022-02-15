package models

import (
	"net/http"

	"github.com/ms-user-portal/app/logging"
)

type User struct {
	Name        *string `json:"name"`
	Location    *string `json:"location"`
	Pan         *string `json:"pan"`
	Address     *string `json:"address"`
	Contact     *string `json:"contact"`
	Sex         *string `json:"sex"`
	Nationality *string `json:"nationality"`
	UserName    *string `json:"userName"`
	Password    *string `json:"password"`
}

type LoginDTO struct {
	UserName *string `json:"userName"`
	Password *string `json:"password"`
}

func (user *User) Validate() *Error {
	lw := logging.LogForFunc()

	if user.Name == nil || len(*user.Name) <= 2 {
		lw.Warn("invalid name")
		return NewError(http.StatusUnprocessableEntity, "invalid name")
	}
	if user.Location == nil || len(*user.Location) <= 3 {
		lw.Warn("invalid location")
		return NewError(http.StatusUnprocessableEntity, "invalid location")
	}
	if user.Pan == nil || len(*user.Pan) != 10 {
		lw.Warn("invalid pan")
		return NewError(http.StatusUnprocessableEntity, "invalid Pan")
	}
	if user.Sex == nil || len(*user.Sex) != 1 {
		lw.Warn("invalid input for sex")
		return NewError(http.StatusUnprocessableEntity, "invalid input for sex")
	}
	if user.UserName == nil || len(*user.UserName) <= 5 {
		lw.Warn("username must be greater than 5 characters")
		return NewError(http.StatusUnprocessableEntity, "username must be greater than 5 characters")
	}
	if user.Password == nil || len(*user.Password) <= 7 {
		lw.Warn("password must be greater than 7 characters")
		return NewError(http.StatusUnprocessableEntity, "password must be greater than 7 characters")
	}
	return nil
}

func (login *LoginDTO) Validate() *Error {
	lw := logging.LogForFunc()
	if login.UserName == nil || len(*login.UserName) == 0 {
		lw.Warn("invalid username")
		return NewError(http.StatusUnprocessableEntity, "invalid username")
	}
	if login.Password == nil || len(*login.Password) == 0 {
		lw.Warn("invalid password")
		return NewError(http.StatusUnprocessableEntity, "invalid password")
	}
	return nil
}
