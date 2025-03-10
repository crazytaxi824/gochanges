package iter_test

import (
	"reflect"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera"}

	// Seq2
	for i, v := range slices.All(names) {
		t.Log(i, v)
	}

	// Seq
	for v := range slices.Values(names) {
		t.Log(v)
	}

	// Seq2
	for k, v := range slices.Backward(names) {
		t.Log(k, v)
	}
}

// iter.Seq
func seq(fn func(int) bool) {
	for i := 0; i < 10; i += 2 {
		fn(i)
	}
}

// 按照一定规则创建 slice
func TestSliceCollect(t *testing.T) {
	s := slices.Collect(seq)
	t.Log(s)
}

// 按照一定规则 append slice
func TestAppendSeq(t *testing.T) {
	s := slices.AppendSeq([]int{1, 2}, seq)
	t.Log(s)
}

func TestAppendSeq2(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := slices.AppendSeq(a, slices.Values(b))
	t.Log(a, b)
	t.Log(c)

	d := slices.Concat(a, b)
	t.Log(a, b)
	t.Log(d)
}

func TestSorted(t *testing.T) {
	a := []int{4, 3, 6, 4, 2, 1}
	b := slices.Clone(a)
	t.Logf("a[0] %p, b[0] %p", &a[0], &b[0])

	slices.Sort(a)
	t.Log("a sort:", a)
	t.Log("b:", b)

	t.Log("b sorted:", slices.Sorted(slices.Values(b)))
	t.Log("b:", b)
}

func TestClip(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(reflect.TypeOf(a))

	s := a[:4:10]
	clip := slices.Clip(s)
	t.Log(cap(s))
	t.Log(cap(clip))
}

func TestChunk(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for v := range slices.Chunk(a, 2) {
		t.Log(v)
	}
	for v := range slices.Chunk(a, 3) {
		t.Log(v)
	}
}
