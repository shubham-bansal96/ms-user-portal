package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/config"
	"github.com/ms-user-portal/app/controllers"
	"github.com/ms-user-portal/app/database"
	"github.com/ms-user-portal/app/middleware"
	"github.com/ms-user-portal/app/services"
)

func Initialize(router *gin.Engine) {

	ctrl := controllers.NewBaseController(services.NewUserService(database.DBConnection), middleware.NewMiddleware())
	appRouter := router.Group(config.Config.MSName)
	appRouter.GET(testEndPoint, ctrl.Ping)
	appRouter.POST(createAccountEndPoint, ctrl.HandleCreateAccount)
	appRouter.POST(loginEndPoint, ctrl.HandleLogin)

	appRouter.Use(middleware.Authorization(ctrl.Middleware))
	appRouter.PATCH(updateEndPoint, ctrl.HandleUpdateUser)
	appRouter.DELETE(deleteEndPoint, ctrl.HandleDeleteAccount)
	appRouter.POST(logoutEndPoint, ctrl.HandleLogout)
}
