package console

import (
	"fmt"
	"github.com/ccb1900/tinyweb/api"
	"github.com/ccb1900/tinyweb/cache"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/db"
	"github.com/ccb1900/tinyweb/flake"
	"github.com/ccb1900/tinyweb/helper"
	log2 "github.com/ccb1900/tinyweb/log"
	gin2 "github.com/ccb1900/tinyweb/middleware/gin"
	"github.com/ccb1900/tinyweb/redis"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

type App struct {
	Engine *gin.Engine
	CliOption *CliOption
	Route func()
}


var app *App

func GetApp() *App  {
	return app
}

func NewApp(f func(),o *CliOption) *App  {
	app = &App{
		Route: f,
		CliOption: o,
	}
	app.beforeEngineStart()
	if config.Get("app.env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app.Engine = gin.New()
	app.afterEngineStart()
	return app
}

func (a *App) Run()  {
	err := a.Engine.Run(getAddr())
	if err != nil {
		debug.PrintStack()
		log2.Fatal(a.CliOption.Config,"not found")
		return
	}
}
func getAddr() string  {
	return fmt.Sprintf("%s:%s",config.Get("app.host"),config.Get("app.port"))
}

func (a *App) beforeEngineStart()  {
	if !helper.IsFileExist(a.CliOption.Config) {
		debug.PrintStack()
		log2.Fatal(a.CliOption.Config,"not found")
	}
	// 加载配置
	config.LoadConfig(a.CliOption.Config)
	// 雪花初始化
	flake.Init()
	// 日志初始化
	log2.Init()
	// 初始化数据库
	db.Init()
	// 初始化redis
	redis.Init()
	// 缓存初始化
	cache.Init()
}

func (a *App)middleware()  {
	app.Engine.Use(gin.Recovery())
	app.Engine.Use(gin.Logger())
	app.Engine.Use(gin2.Cors)
}

func (a *App)route()  {
	// 路径不匹配
	app.Engine.NoRoute(func(c *gin.Context) {
		api.NotFound(c)
	})
	// 方法不匹配
	app.Engine.NoMethod(func(c *gin.Context) {
		api.MethodNotMatch(c)
	})
	// 上传文件路径
	if config.Get("upload.driver") == "local" {
		app.Engine.Static("resources", config.Get("upload.local.path"))
	}
	app.Route()
}
func (a *App) afterEngineStart()  {
	a.middleware()
	a.route()
}
