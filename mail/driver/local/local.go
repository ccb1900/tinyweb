package local

import (
	"github.com/ccb1900/tinyweb/log"
	"github.com/ccb1900/tinyweb/mail/contract"
)

type Local struct {
	
}

func (l *Local)Send(email,content string)  {
	go func() {
		log.Debug(email,content)
	}()
}

func New() contract.IMail  {
	return &Local{}
}
