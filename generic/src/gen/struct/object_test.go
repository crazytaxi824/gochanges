package struct_test

import "testing"

type MyStruct[K comparable, V any] struct {
	Key   K
	Value V
}

// object method, 在这里会自动推导出 object MyStruct 中的约束为 K comparable, V any
func (m MyStruct[K, V]) Map() map[K]V {
	return map[K]V{
		m.Key: m.Value,
	}
}

func genMap[K comparable, V any](m MyStruct[K, V]) map[K]V {
	return m.Map()
}

func TestMap(t *testing.T) {
	m := MyStruct[string, int]{
		Key:   "one",
		Value: 1,
	}

	t.Log(genMap(m))
}
