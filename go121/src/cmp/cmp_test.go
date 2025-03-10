package cmp_test

import (
	"cmp"
	"testing"
)

// less: -1; equal: 0; greater: 1
// func Compare[T Ordered](x, y T) int

// less: true; equal: false; greater: false
// func Less[T Ordered](x, y T) bool

func TestCmp(t *testing.T) {
	t.Log(cmp.Compare(1, 6)) // -1
	t.Log(cmp.Compare(1, 1)) // 0
	t.Log(cmp.Compare(6, 1)) // 1

	t.Log(cmp.Less(1, 2)) // true
	t.Log(cmp.Less(1, 1)) // false
	t.Log(cmp.Less(2, 1)) // false
}
