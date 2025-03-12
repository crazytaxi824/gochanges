package func_test

import "testing"

// Type Param
// 该函数可以接受一个底层为 int OR int64 类型的参数传入, 返回相同的类型.
func echoTypeParam[T interface{ ~int | ~int64 }](t T) T {
	return t
}

// 小技巧: 返回零值
func zeroValueGeneric[T interface{ ~int | ~string | ~bool }]() T {
	return *new(T)
}

func TestZeroVale(t *testing.T) {
	t.Log(zeroValueGeneric[int]())    // 0
	t.Log(zeroValueGeneric[bool]())   // false
	t.Log(zeroValueGeneric[string]()) // ""
}

// 复杂的 Type Param, 该 type params 中, 我们需要定义一个约束 map[comparable]any
func makeMap[M map[K]V, K comparable, V any](k K, v V) M {
	return map[K]V{
		k: v,
	}
}

func makeSlice[S ~[]E, E any](e E) S {
	return []E{e}
}
