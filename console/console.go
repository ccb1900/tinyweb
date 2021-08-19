package console

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func CliApp()  {
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
					start(ctx)
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			start(ctx)
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		log.Println(err)
		return
	}
}

func start(ctx *cli.Context)  {
	log.Println("sssss")
	ctx.App.Usage = "德玛西亚"
	ctx.App.Name = "德玛西亚"
}
