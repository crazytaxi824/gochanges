package waitgroup

import (
	"sync"
	"testing"
)

// 传统方式
func TestWaitGroup(t *testing.T) {
	var w sync.WaitGroup

	// 更灵活
	for i := range 3 {
		w.Add(1)
		go func(n int) {
			defer w.Done()
			t.Log(n)
		}(i)
	}
	w.Wait()
}

// WaitGroup.Go() 新方法
func TestWaitGroup2(t *testing.T) {
	var w sync.WaitGroup

	for i := range 3 {
		// 更方便
		w.Go(func() {
			t.Log(i)
		})
	}
	w.Wait()
}
