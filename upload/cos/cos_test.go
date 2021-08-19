package cos

import (
	"github.com/ccb1900/tinyweb/config"
	"testing"
)

func TestStart(t *testing.T) {
	config.LoadConfig("../../config/app.dev.toml")
	start()
}
