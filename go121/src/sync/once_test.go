package sync_test

import (
	"sync"
	"testing"
)

// sync.Once
// func (o *Once) Do(f func())

// eg:
func TestLazyInit(t *testing.T) {
	type lazyCache struct {
		M map[string]int
		sync.Once
	}

	var m lazyCache

	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Once.Do(func() {
				m.M = make(map[string]int)
				t.Logf("goroutine %d run once", n)
			})
		}(i)
	}
	wg.Wait()
}

// OnceFunc() 中的 fn 只会运行一次
func TestOnceFunc(t *testing.T) {
	var lazyCache map[string]int

	onceFn := sync.OnceFunc(func() {
		lazyCache = make(map[string]int)
		t.Logf("lazyCache init once")
	})

	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			onceFn()
		}()
	}
	wg.Wait()

	lazyCache["one"] = 1
}

// OnceValue 相当于 OnceFunc with a return
// OnceValue 中的 fn 只会运行一次, 但是每次运行 fn 都会得到同一个 return.
func TestOnceValue(t *testing.T) {
	var lazyCache map[string]int

	onceFn := sync.OnceValue(func() bool {
		lazyCache = make(map[string]int)
		t.Logf("lazyCache init once")
		return true
	})

	var wg = new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			t.Log(n, onceFn())
		}(i)
	}
	wg.Wait()

	lazyCache["one"] = 1
}
