package config

import (
	"sync"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("./app.example.toml")
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			t.Run("app.name", func(t *testing.T) {
				t.Log(i,Get("app.name")=="app")
				Set("app.name","demo")
			})
		}(i)
	}
	wg.Wait()
}
