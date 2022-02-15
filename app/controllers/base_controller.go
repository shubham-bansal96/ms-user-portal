package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/logging"
	"github.com/ms-user-portal/app/middleware"
	"github.com/ms-user-portal/app/models"
	"github.com/ms-user-portal/app/services"
	"github.com/ms-user-portal/app/utils"
)

type BaseController struct {
	UserService services.IUser
	Middleware  middleware.IToken
}

func NewBaseController(userService services.IUser, middleware middleware.IToken) *BaseController {
	return &BaseController{
		UserService: userService,
		Middleware:  middleware,
	}
}

// Ping => Get ping from microservice
func (ctrl *BaseController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"test": "ok"})
}

// HandleCreateAccount => create user account
func (ctrl *BaseController) HandleCreateAccount(ctx *gin.Context) {
	lw := logging.LogForFunc()
	var reqObj *models.User

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		lw.WithField("error", "error occured while unmarshelling JSON request").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, models.NewError(http.StatusBadRequest, "error occured while unmarshelling"))
		return
	}

	if err := reqObj.Validate(); err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	_, err := ctrl.UserService.CreateAccount(ctx, reqObj)
	if err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	lw.Info("account created succssfully")
	utils.RendorJson(ctx, "account created successfully", http.StatusCreated, nil)
	return
}

// HandleLogin => Login to user account
func (ctrl *BaseController) HandleLogin(ctx *gin.Context) {
	lw := logging.LogForFunc()
	var reqObj *models.LoginDTO

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		lw.WithField("error", "error occured while unmarshelling JSON request").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, models.NewError(http.StatusBadRequest, "error occured while unmarshelling"))
		return
	}

	if err := reqObj.Validate(); err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	userID, err := ctrl.UserService.Login(ctx, reqObj)
	if err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	token, err := ctrl.Middleware.GenerateToken(ctx, *userID)
	if err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}
	utils.RendorJson(ctx, token, http.StatusOK, err)
	return
}

// HandleUpdateUser => update user account
func (ctrl *BaseController) HandleUpdateUser(ctx *gin.Context) {
	lw := logging.LogForFunc()
	var reqObj *models.User

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		lw.WithField("error", "error occured while unmarshelling JSON request").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, models.NewError(http.StatusBadRequest, "error occured while unmarshelling"))
		return
	}

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lw.WithField("error", "invalid userid").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, models.NewError(http.StatusBadRequest, "invalid userid"))
		return
	}

	if err := ctrl.UserService.UpdateAccount(ctx, reqObj, userID); err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	utils.RendorJson(ctx, "account updated successfully", http.StatusOK, nil)
	return
}

// HandleUpdateUser => delete user account
func (ctrl *BaseController) HandleDeleteAccount(ctx *gin.Context) {
	lw := logging.LogForFunc()
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lw.WithField("error", "invalid userid").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, models.NewError(http.StatusBadRequest, "invalid userid"))
		return
	}

	if err := ctrl.UserService.DeleteAccount(ctx, userID); err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	utils.RendorJson(ctx, "account deleted successfully", http.StatusOK, nil)
	return
}

// HandleLogout => logout user
func (ctrl *BaseController) HandleLogout(ctx *gin.Context) {
	claim := &models.Claims{}
	claim.DeleteLoginCookies(ctx, "")
	ctx.JSON(http.StatusNoContent, nil)
	return
}
