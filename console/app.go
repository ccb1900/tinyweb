package console

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
	Route func()
}


var app *App

func GetApp() *App  {
	return app
}

func SetApp(f func())  {
	if config.Get("app.env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app = &App{
		Engine: gin.New(),
		Route: f,
	}
}
