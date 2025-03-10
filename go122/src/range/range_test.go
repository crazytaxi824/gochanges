package range_test

import (
	"sync"
	"testing"
)

// 新用法
func TestRange(t *testing.T) {
	for i := range 3 {
		t.Log(i)
	}
}

// Fixing For Loops var in Go 1.22
// https://go.dev/blog/loopvar-preview
func TestRangeVar(t *testing.T) {
	sl := []int{11, 12, 13, 14, 15}
	var ri []*int
	var rv []*int
	for i, v := range sl {
		ri = append(ri, &i)
		rv = append(rv, &v)
	}

	for i := range sl {
		// <= v1.21 结果为:
		// 4 15
		// 4 15
		// 4 15
		// 4 15
		// 4 15
		// >= v1.22 结果为:
		// 0 11
		// 1 12
		// 2 13
		// 3 14
		// 4 15
		t.Log(*ri[i], *rv[i])
	}
}

// 闭包函数
func TestClosureFn(t *testing.T) {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for k, v := range sl {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// <= v1.21 结果为:
			// 4 15
			// 4 15
			// 4 15
			// 4 15
			// 4 15
			// >= v1.22 结果为:
			// 4 15
			// 2 13
			// 3 14
			// 0 11
			// 1 12
			t.Log(k, v)
		}()
	}
	wg.Wait()
}
