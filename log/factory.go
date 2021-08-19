package log

import (
	"github.com/ccb1900/tinyweb/log/driver"
	"log"
)

func Factory(name string) ILog  {
	switch name {
	case "default":
		return &driver.Std{
			Name: name,
			Log: log.Default(),
		}
	case "std":
		return &driver.Std{
			Name: name,
			Log: log.Default(),
		}
	case "zap":
		return &driver.Zap{

		}
	default:
		panic("日志驱动配置错误")
	}

}
