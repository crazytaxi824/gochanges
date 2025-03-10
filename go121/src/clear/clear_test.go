package clear_test

import "testing"

func TestClearSlice(t *testing.T) {
	a := [10]int{1, 2, 3, 4, 5, 6}
	b := a[:]
	t.Log(a) // [1 2 3 4 5 6 0 0 0 0]
	t.Log(b) // [1 2 3 4 5 6 0 0 0 0]

	clear(b)
	t.Log(b) // [0 0 0 0 0 0 0 0 0 0]
}

func TestClearMap(t *testing.T) {
	m := map[string]any{
		"foo": 1,
	}

	clear(m)
	t.Log(m) // delete all kv
}
