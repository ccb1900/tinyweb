package console

import "github.com/gin-gonic/gin"

type App struct {
	Engine *gin.Engine
	Route func()
}


var app *App

func GetApp() *App  {
	return app
}

func SetApp(f func())  {
	app = &App{
		Engine: gin.New(),
		Route: f,
	}
}
