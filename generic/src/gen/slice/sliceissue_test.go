package slice_test

import (
	"slices"
	"strings"
	"testing"
)

// 以下两种情况带来的区别
func clone1[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

func clone2[E any](s []E) []E {
	return slices.Clone(s)
}

// 实现
type myClone []string

func (m myClone) String() string {
	return strings.Join(m, "+")
}

func TestClone(t *testing.T) {
	var m myClone

	m1 := clone1(m)
	m2 := clone2(m)

	m1.String()
	// m2.String() // ERROR: 没有该方法

	// print Type
	t.Logf("%T", m)  // myClone
	t.Logf("%T", m1) // myClone
	t.Logf("%T", m2) // []string, 所以没有 String() 方法了
}
