package rest

import (
	"github.com/gin-gonic/gin"
	logger "sales-go/helpers/logging"
)


func ResponseData(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"status": status,
		"data"  : data,
	})
}

func ResponseError(ctx *gin.Context, status int, err error) {
	logger.Errorf(err, ctx.Request)

	ctx.JSON(status, map[string]interface{}{
		"status": status,
		"error":  err.Error(),
	})
}