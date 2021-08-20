package console

import (
	"fmt"
	"github.com/ccb1900/tinyweb/api"
	"github.com/ccb1900/tinyweb/cache"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/db"
	"github.com/ccb1900/tinyweb/helper"
	gin2 "github.com/ccb1900/tinyweb/middleware/gin"
	"github.com/ccb1900/tinyweb/redis"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime/debug"
)

func CliApp(f func())  {
	err := (&cli.App{
		Name: "ly ui",
		Description: "a web app",
		Usage: "nothing",
		// 全局选项
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "c",
				Value:        "./cmd/res/app.dev.toml",
				Usage: "load config",
			},
		},
		// 对应于不同的操作
		Commands: []*cli.Command{
			{
				Name: "start",
				Action: func(ctx *cli.Context) error {
					start(ctx,f)
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			start(ctx,f)
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		log.Println(err)
		return
	}
}

// 应用入口
func start(ctx *cli.Context,f func())  {
	beforeAppStart(ctx)
	SetApp(f)
	afterAppStart()
	err := app.Engine.Run(getAddr())

	if err != nil {
		debug.PrintStack()
		log.Fatal(ctx.String("c"),"not found")
		return 
	}
}

func getAddr() string  {
	return fmt.Sprintf("%s:%s",config.Get("app.host"),config.Get("app.port"))
}

func beforeAppStart(ctx *cli.Context)  {
	if !helper.IsFileExist(ctx.String("c")) {
		debug.PrintStack()
		log.Fatal(ctx.String("c"),"not found")
	}
	// 加载配置
	config.LoadConfig(ctx.String("c"))
	// 初始化数据库
	db.Init()
	// 初始化redis
	redis.Init()
	// 缓存初始化
	cache.Init()

}

func middleware()  {
	app.Engine.Use(gin.Recovery())
	app.Engine.Use(gin.Logger())
	app.Engine.Use(gin2.Cors)
}

func route()  {
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
func afterAppStart()  {
	if config.Get("app.env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	middleware()
	route()
}
