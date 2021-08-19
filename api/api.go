package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context,obj interface{})  {
	ctx.JSON(http.StatusOK,obj)
}

func Fail(ctx *gin.Context,message string)  {
	ctx.JSON(http.StatusBadRequest,gin.H{
		"err": message,
	})
}
func FailValidation(ctx *gin.Context,message string)  {
	ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		"err": message,
	})
}
func NotAuth(ctx *gin.Context,message string)  {
	ctx.JSON(http.StatusUnauthorized,gin.H{
		"err": message,
	})
}
