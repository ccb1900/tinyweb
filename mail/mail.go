package mail

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/mail/contract"
	"github.com/ccb1900/tinyweb/mail/driver/local"
)

var m contract.IMail

func Init()  {
	switch config.Get("mail") {
	case "local":
		m = local.New()
	}
}

func Send(email,content string)  {
	m.Send(email,content)
}
