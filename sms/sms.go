package sms

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/sms/contract"
	"github.com/ccb1900/tinyweb/sms/driver/local"
)

var s contract.ISms
func Init()  {
	switch config.Get("sms.driver") {
	case "local":
		s = local.New()
	default:
		s = local.New()
	}
}

func Send(mobile string, content string) {
	s.Send(mobile,content)
}

