package iter_test

import (
	"maps"
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	m1 := map[string][]int{
		"one": {1},
		"two": {2},
	}
	m2 := map[string][]int{
		"one":   {10},
		"three": {3},
	}

	// force combine, elements shallow copy
	maps.Insert(m1, maps.All(m2))
	t.Log(m1, m2)

	// m2 elements 改变会改变 m1
	m2["one"][0] = 99
	t.Log(m1, m2)

	// m2 整体改变不会影响 m1
	m2["one"] = []int{100}
	t.Log(m1, m2)
}

func TestCopy(t *testing.T) {
	m1 := map[string][]int{
		"one": {1},
		"two": {2},
	}
	m2 := map[string][]int{
		"one":   {10},
		"three": {3},
	}

	// force combine, elements shallow copy
	maps.Copy(m1, m2)
	t.Log(m1, m2)

	// m2 elements 改变会改变 m1
	m2["one"][0] = 99
	t.Log(m1, m2)

	// m2 整体改变不会影响 m1
	m2["one"] = []int{100}
	t.Log(m1, m2)
}

func TestClone(t *testing.T) {
	m := map[string][]int{
		"one": {1},
		"two": {2},
	}

	// elements shallow copy
	mclone := maps.Clone(m)
	m["one"][0] = 99
	t.Log(m, mclone)

	mclone["three"] = []int{3}
	t.Log(m, mclone)
}

func TestMapCollect(t *testing.T) {
	s1 := []string{"zero", "one", "two", "three"}
	m1 := maps.Collect(slices.All(s1))
	t.Log("m1 is:", m1) // 输出: map[0:zero 1:one 2:two]
}
