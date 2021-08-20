package log

import (
	"github.com/ccb1900/tinyweb/config"
	"sync"
	"testing"
)

func TestDebug(t *testing.T) {
	config.LoadConfig("../config/app.example.toml")
	Init()
	call()
	c(100, func(i int) {
		Debug(i,"debug")
	})
}
func call()  {
	Debug("test")
}
func c(num int,task func(int))  {
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			task(i)
		}(i)
	}
	wg.Wait()
}
