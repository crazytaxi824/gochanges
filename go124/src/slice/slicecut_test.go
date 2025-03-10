package slice_test

import (
	"slices"
	"testing"
)

func TestSliceGrow(t *testing.T) {
	d := []byte("abc")
	t.Log(len(d), cap(d))

	d = slices.Grow(d, 10)
	t.Log(len(d), cap(d))

	// 不能超过 cap
	t.Log(d[3:])
	t.Log(d[3:cap(d)])
	// t.Log(d[3 : cap(d)+1]) // Panic out of range
	// t.Log(d[11]) // Panic out of range

	// 不能超过 cap
	t.Log(d[3:][:cap(d)-3])
	// t.Log(d[3:][:cap(d)-2]) // Panic out of range
}

func TestSliceCopy(t *testing.T) {
	d := make([]byte, 0, 16)
	d = append(d, 'a', 'b', 'c')
	t.Log(d)          // [97 98 99]
	t.Log(d[3:][:13]) // [0 0 0 0 0 0 0 0 0 0 0 0 0]

	// 并没有改变 d, 但是改变了底层 array 的数据.
	copy(d[3:][:13], "xyz")
	t.Log(d)          // [97 98 99]
	t.Log(d[3:][:13]) // [120 121 122 0 0 0 0 0 0 0 0 0 0]

	// 两种写法都可以
	copy(d[6:16], "foo")
	t.Log(d)          // [97 98 99]
	t.Log(d[3:][:13]) // [120 121 122 102 111 111 0 0 0 0 0 0 0]
}
