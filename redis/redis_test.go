package redis

import (
	"github.com/ccb1900/tinyweb/config"
	"testing"
)

func TestClient(t *testing.T) {
	config.LoadConfig("../config/app.dev.toml")
	client()
}
