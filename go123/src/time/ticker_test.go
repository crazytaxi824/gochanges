package time_test

import (
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	tr := time.NewTicker(time.Second)

	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()

		// ticker 循环执行 C <-
		for k := range tr.C {
			t.Log(k)
		}
	}()

	wg.Wait()
}
