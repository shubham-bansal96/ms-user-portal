package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/models"
)

func RendorJson(ctx *gin.Context, data interface{}, statusCode int, err *models.Error) {
	responseDTO := &models.ResponseDTO{
		Data:  data,
		Error: err,
	}

	ctx.JSON(statusCode, responseDTO)
}
