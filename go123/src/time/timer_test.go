package time_test

import (
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	tr := time.NewTimer(time.Second)

	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()

		// timer 时间到后执行一次 C <-
		for k := range tr.C {
			t.Log(k)
			return
		}
	}()

	wg.Wait()
}

func TestTimerStop(t *testing.T) {
	tr := time.NewTimer(time.Second)

	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()

		// timer 时间到后执行一次 C <-
		for k := range tr.C {
			t.Log(k)
			return
		}
	}()

	// stop 之后会停止向 C <-, 但是不会 close channel
	tr.Stop()
	t.Log("time stopped")

	t.Log("sleep")
	time.Sleep(3 * time.Second)

	t.Log("reset")
	tr.Reset(2 * time.Second)

	wg.Wait()
}
