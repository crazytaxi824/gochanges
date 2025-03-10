// Range over Function
// for range 中的 func 需要符合以下几个类型:
// func(func() bool)
// func(func(V) bool)  -- iter.Seq
// func(func(K, V) bool)  -- iter.Seq2

package iter_test

import (
	"testing"
)

// func(func() bool)
func loop(count int) func(func() bool) {
	return func(fn func() bool) {
		for range count {
			fn()
		}
	}
}

// 这里可以看出 Slice2 返回的就是一个 iter.Seq2 类型
func sli[V any](s []V) func(func(V) bool) {
	return func(fn func(V) bool) {
		for _, v := range s {
			// 将遍历的数据传入 func(func(v) bool) 中。
			fn(v)
		}
	}
}

// 这里可以看出 sli2 返回的就是一个 iter.Seq2 类型
func sli2[V any](s []V) func(func(int, V) bool) {
	return func(fn func(int, V) bool) {
		for k, v := range s {
			// 将遍历的数据传入 func(func(k,v) bool) 中。
			fn(k, v)
		}
	}
}

// iter.Seq2 类型
func mapN[K comparable, V any](m map[K]V) func(func(K, V) bool) {
	return func(fn func(K, V) bool) {
		for k, v := range m {
			fn(k, v)
		}
	}
}

// for range func(func(k,v) bool) 也主要是为了 iter 服务的。
func TestForRange(t *testing.T) {
	// func(func() bool)
	for range loop(3) {
		t.Log("loop")
	}

	// func(func(V) bool)  -- iter.Seq
	for v := range sli([]int{3, 2, 1}) {
		t.Log(v)
	}

	// func(func(K, V) bool)  -- iter.Seq2
	for i, v := range sli2([]int{7, 8, 9}) {
		t.Log(i, v)
	}

	// func(func(K, V) bool)  -- iter.Seq2
	for k, v := range mapN(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}) {
		t.Log(k, v)
	}
}
