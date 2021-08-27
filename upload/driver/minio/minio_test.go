package minio

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/log"
	"testing"
)

func TestNew(t *testing.T) {
	config.LoadConfig("../../../config/app.dev.toml")
	log.Init()
	t.Log(New())
	u :=New()
	err := u.Upload(nil, "")
	if err != nil {
		t.Log(err)
		return
	}
}
