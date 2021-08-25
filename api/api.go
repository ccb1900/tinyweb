package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context,obj interface{})  {
	ctx.JSON(http.StatusOK,obj)
}

func Ok(c *gin.Context)  {
	Success(c,gin.H{
		"message": "操作成功",
	})
}

func Created(ctx *gin.Context,obj interface{})  {
	ctx.JSON(http.StatusCreated,obj)
}

func Fail(ctx *gin.Context,err error)  {
	ctx.JSON(http.StatusBadRequest,gin.H{
		"err": err.Error(),
	})
}
func FailValidation(ctx *gin.Context,err error)  {
	ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		"err": err.Error(),
	})
}
func NotAuth(ctx *gin.Context,message string)  {
	ctx.JSON(http.StatusUnauthorized,gin.H{
		"err": message,
	})
}
func NotFound(ctx *gin.Context)  {
	ctx.JSON(http.StatusNotFound,gin.H{
		"err": "404",
	})
}

func MethodNotMatch(ctx *gin.Context)  {
	ctx.JSON(http.StatusMethodNotAllowed,gin.H{
		"err": "405",
	})
}
