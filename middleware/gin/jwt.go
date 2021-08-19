package gin

import (
	"github.com/ccb1900/tinyweb/api"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Header struct {
	Authorization string `json:"Authorization" binding:"required"`
}

func Jwt(ctx *gin.Context)  {
	if config.Get("app.env") == "dev" {
		ctx.Set("username","admin")
		ctx.Next()
		return
	}
	h := Header{}
	err := ctx.ShouldBindHeader(&h)
	if err != nil {
		ctx.Abort()
		log.Println(err.Error())
		api.NotAuth(ctx,"头错误")
		return
	}

	pieces := strings.Split(strings.Trim(h.Authorization," ")," ")
	log.Println("pieces",len(pieces),pieces)

	if len(pieces) == 0 {
		ctx.Abort()
		api.NotAuth(ctx,"令牌不能为空")
		return
	}
	if strings.ToLower(pieces[0]) != "bearer" {
		ctx.Abort()
		api.NotAuth(ctx,"授权类型错误")
		return
	}

	claims, err := jwt.ParseJwtToken(pieces[len(pieces)-1])
	if err != nil {
		ctx.Abort()
		api.NotAuth(ctx,"无效的令牌")
		return
	}

	ctx.Set("username",claims.UserName)
	ctx.Set("id",claims.ID)
	ctx.Next()
}
