package flake

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/flake/contract"
	"github.com/ccb1900/tinyweb/flake/driver/sony"
)

var fl contract.IFlake


func NextId() uint64  {
	return fl.NextId()
}

func Init()  {
	switch config.Get("flake.driver") {
	case "sony":
		fl = sony.New()
	}
}
