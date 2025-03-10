package slog_test

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	a := [10]int{}
	b := a[:]
	b[0] = 1
	t.Log(a, reflect.TypeOf(a))
	t.Log(b, reflect.TypeOf(b))
}

func BenchmarkSlice(b *testing.B) {
	for b.Loop() {
		// code
		_ = make([]int, 100)
	}
	b.ReportAllocs() // go test -benchmem
}

func BenchmarkArray(b *testing.B) {
	for b.Loop() {
		// code
		a := [100]int{}
		_ = a[:]
	}
	b.ReportAllocs() // go test -benchmem
}
