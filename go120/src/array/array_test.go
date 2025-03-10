package array_test

import "testing"

func TestArray(t *testing.T) {
	var s = []int{1, 2, 3}
	var p = (*[3]int)(s) // reference
	var a = [3]int(s)    // copy

	// 0x1400001e168 0x1400001e168 0x1400001e180
	t.Logf("%p %p %p", &s[0], &p[0], &a[0])

	// a is a copy of s
	a[1] = a[1] + 100
	t.Log(a, s) // [1 102 3] [1 2 3]

	// p is a ref of s
	p[1] = p[1] + 100
	t.Log(p, s) // &[1 102 3] [1 102 3]
}
