package iter_test

import (
	"maps"
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 10,
	}

	// overwrite
	maps.Insert(m2, maps.All(m1))
	t.Log("m2 is:", m2)
}

func TestClone(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	mclone := maps.Clone(m)
	mclone["three"] = 3
	t.Log(m, mclone)
}

func TestCopy(t *testing.T) {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 10,
	}

	// overwrite
	maps.Copy(m1, m2)
	t.Log(m1)
}

func TestMapCollect(t *testing.T) {
	s1 := []string{"zero", "one", "two", "three"}
	m1 := maps.Collect(slices.All(s1))
	t.Log("m1 is:", m1)
}
