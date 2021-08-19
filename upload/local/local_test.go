package local

import (
	"github.com/ccb1900/tinyweb/config"
	"testing"
)

func TestUpload(t *testing.T) {
	config.LoadConfig("../../config/app.dev.toml")
	upload()
}
