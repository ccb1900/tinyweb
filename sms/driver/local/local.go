package local

import (
	"github.com/ccb1900/tinyweb/log"
	"github.com/ccb1900/tinyweb/sms/contract"
)

type Local struct {
	
}

func (l *Local) Send(mobile string, content string) {
	go func() {
		log.Debug(mobile,content)
	}()
}

func New() contract.ISms {
	return &Local{}
}
