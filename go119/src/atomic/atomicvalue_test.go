package atomic_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomicValue(t *testing.T) {
	var v atomic.Value

	v.Store("a")
	t.Log(v.Load()) // a

	v.Store("b")
	t.Log(v.Load()) // b

	t.Log(v.CompareAndSwap("b", "c")) // true
	t.Log(v.Load())                   // c

	t.Log(v.CompareAndSwap("b", "d")) // false
	t.Log(v.Load())                   // c
}

func TestAtomicAdd(t *testing.T) {
	var a int64

	t.Log(atomic.AddInt64(&a, 3))
	t.Log(a)
}

func TestAtomicAdd2(t *testing.T) {
	var a int64

	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			t.Log(n, atomic.AddInt64(&a, n))
		}(int64(i))
	}

	wg.Wait()
	t.Log(a)
}
