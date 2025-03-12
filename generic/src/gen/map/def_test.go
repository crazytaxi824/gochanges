package map_test

// define map generic
type myMap interface {
	map[string]any | *map[string]any | []map[string]any | []*map[string]any
}

type myMap2 interface {
	~map[string]any | ~*map[string]any | ~[]map[string]any | ~[]*map[string]any
}

type myMap3[K comparable, V any] interface {
	map[K]V | *map[K]V | []map[K]V | []*map[K]V
}

type MyMapStruct[K comparable, V any] struct {
	m map[K]V
}

// func funcTypeParamMap[M map[K]V, K, V any](M) {} // ERROR: invalid map key type K
func funcTypeParamMap[M map[K]V, K comparable, V any](k K, v V) M {
	return map[K]V{k: v}
}
