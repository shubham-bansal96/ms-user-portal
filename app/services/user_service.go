package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ms-user-portal/app/database"
	"github.com/ms-user-portal/app/logging"
	"github.com/ms-user-portal/app/models"
)

type IUser interface {
	CreateAccount(ctx context.Context, user *models.User) (*int, *models.Error)
	Login(ctx context.Context, loginDTO *models.LoginDTO) (*int, *models.Error)
	UpdateAccount(ctx context.Context, user *models.User, userId int) *models.Error
	DeleteAccount(ctx context.Context, userId int) *models.Error
}

type User struct {
	DBConnection database.IConnection
}

func NewUserService(dbObj database.IConnection) IUser {
	return &User{
		DBConnection: dbObj,
	}
}

func (user *User) CreateAccount(ctx context.Context, userObj *models.User) (*int, *models.Error) {

	lw := logging.LogForFunc()

	var (
		message string
		userID  int
	)

	_, err := user.DBConnection.Exec(ctx, "USP_CreateUserAccount",
		sql.Named("Name", userObj.Name),
		sql.Named("Location", userObj.Location),
		sql.Named("Pan", userObj.Pan),
		sql.Named("Address", userObj.Address),
		sql.Named("Contact", userObj.Contact),
		sql.Named("Sex", string(*userObj.Sex)),
		sql.Named("Nationality", userObj.Nationality),
		sql.Named("UserName", userObj.UserName),
		sql.Named("Password", userObj.Password),
		sql.Named("Msg", sql.Out{Dest: &message}),
		sql.Named("UserID", sql.Out{Dest: &userID}))

	if err != nil {
		lw.WithField("error", "error while executing USP_CreateUserAccount").Error(err.Error())
		return nil, models.NewError(http.StatusInternalServerError, "error while executing USP_CreateUserAccount")
	}

	if message != "" {
		lw.WithField("error", "some error occured in USP_CreateUserAccount").Error(message)
		return nil, models.NewError(http.StatusUnprocessableEntity, message)
	}
	return &userID, nil
}

func (user *User) Login(ctx context.Context, loginDTO *models.LoginDTO) (*int, *models.Error) {
	lw := logging.LogForFunc()

	var userID int
	row := user.DBConnection.QueryRow(ctx, "SELECT ID AS userID FROM User_Data WHERE UserName = @username AND Password = @password",
		sql.Named("username", loginDTO.UserName),
		sql.Named("password", loginDTO.Password))

	err := row.Scan(&userID)
	if err == sql.ErrNoRows {
		lw.WithField("warn", "invalid login credentials").Warn(err.Error())
		return nil, models.NewError(http.StatusInternalServerError, "invalid login credentials")
	}
	if err != nil {
		lw.WithField("error", "some error occured while executing query").Error(err.Error())
		return nil, models.NewError(http.StatusInternalServerError, "some error occured while executing query")
	}

	return &userID, nil
}

func (user *User) UpdateAccount(ctx context.Context, userObj *models.User, userId int) *models.Error {
	lw := logging.LogForFunc()

	result, err := user.DBConnection.Exec(ctx, "UPDATE User_Data SET Name= @name, Pan=@pan WHERE id=@userId",
		sql.Named("name", userObj.Name),
		sql.Named("pan", userObj.Pan),
		sql.Named("userId", userId))

	if err != nil {
		lw.WithField("error", "error while executing query").Error(err.Error())
		return models.NewError(http.StatusInternalServerError, "error while executing query")
	}
	rows, err2 := result.RowsAffected()
	if err2 != nil {
		lw.WithField("error", "error while updating account").Error(err2.Error())
		return models.NewError(http.StatusInternalServerError, "error while updating account")
	}
	if rows == 0 {
		lw.WithField("warn", "no data found to update").Warn(err)
		return models.NewError(http.StatusUnprocessableEntity, "no data found to update")
	}
	return nil
}

func (user *User) DeleteAccount(ctx context.Context, userId int) *models.Error {
	lw := logging.LogForFunc()

	result, err := user.DBConnection.Exec(ctx, "DELETE FROM User_Data WHERE id=@userId",
		sql.Named("userId", userId))

	if err != nil {
		lw.WithField("error", "error while executing query").Error(err.Error())
		return models.NewError(http.StatusInternalServerError, "error while executing query")
	}
	rows, err2 := result.RowsAffected()
	if err2 != nil {
		lw.WithField("error", "error while deleting account").Error(err2.Error())
		return models.NewError(http.StatusInternalServerError, "error while deleting account")
	}
	if rows == 0 {
		lw.WithField("warn", "no record found to delete").Warn(err)
		return models.NewError(http.StatusUnprocessableEntity, "no record found to update")
	}
	return nil
}
