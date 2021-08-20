package flake

import (
	"sync"
	"testing"
)

func TestNextId(t *testing.T) {
	Init()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				t.Log(NextId())
			}
		}()
	}
	wg.Wait()

}

func TestInit(t *testing.T) {
	t.Log(len("6439921798267012072"))
	t.Log(len("1424662605326512128"))
}
