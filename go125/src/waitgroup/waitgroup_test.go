package waitgroup

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var w sync.WaitGroup

	for i := range 3 {
		w.Add(1)
		go func() {
			defer w.Done()
			t.Log(i)
		}()
	}
	w.Wait()
}

// WaitGroup.Go() 新方法
func TestWaitGroup2(t *testing.T) {
	var w sync.WaitGroup

	for i := range 3 {
		w.Go(func() {
			t.Log(i)
		})
	}
	w.Wait()
}
