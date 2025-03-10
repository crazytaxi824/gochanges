package map_test

import (
	"maps"
	"testing"
)

func TestMapClone(t *testing.T) {
	m := map[string]any{
		"foo": 1,
		"bar": "two",
	}

	m2 := maps.Clone(m)
	m["foo"] = 99
	t.Log(m, m2) // shallow clone
	// 如果 v 本身是引用类型, 则不会继续深入 clone.
}

func TestMapCopy(t *testing.T) {
	m1 := map[string]any{
		"foo": 1,
		"bar": "two",
	}
	m2 := map[string]any{ // key exist
		"foo": "foo",
	}
	m3 := map[string]any{ // kv > src map
		"kk": 9,
	}

	maps.Copy(m2, m1) // m1 overwrite m2
	t.Log(m2)

	maps.Copy(m3, m1)
	t.Log(m3)
}

func TestMapDelete(t *testing.T) {
	m1 := map[string]int{
		"foo": 1,
		"bar": 99,
	}

	// 按条件删除 KV
	maps.DeleteFunc(m1, func(_ string, v int) bool {
		return v > 10
	})
	t.Log(m1)

	// 按 Key 删除
	delete(m1, "bar")
	t.Log(m1)
}

func TestMapEq(t *testing.T) {
	m1 := map[string]int{
		"foo": 1,
		"bar": 99,
	}

	m2 := map[string]int{
		"bar": 99,
		"foo": 1,
	}
	t.Log(maps.Equal(m1, m2)) // true
}
