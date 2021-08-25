package sms

import (
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/log"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	config.LoadConfig("../config/app.dev.toml")
	log.Init()
	Init()
	Send("15601948562","短信来了啊")
	time.Sleep(2*time.Second)
	fmt.Println("wait...")
}
