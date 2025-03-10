package atomic_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomicInt64(t *testing.T) {
	var a atomic.Int64
	t.Log(a.Load()) // 0

	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			t.Log(n, a.Add(n))
		}(int64(i))
	}
	wg.Wait()
	t.Log(a.Load())
}

func TestAtomicPointer(t *testing.T) {
	var p atomic.Pointer[int]
	t.Log(p.Load())

	a := 99
	p.Store(&a)
	t.Log(p.Load(), *p.Load())
}

func TestAtomicPointer2(t *testing.T) {
	var p atomic.Pointer[func() bool]
	t.Log(p.Load())

	var a = func() bool {
		return true
	}
	p.Store(&a)
	t.Log((*p.Load())())
}
