// 新增 b.Loop() 方法用于 benchmem test
package bench_test

import (
	"testing"
)

// 之前的做法
func BenchmarkBN(b *testing.B) {
	for range b.N {
		// code
	}
	for i := 0; i < b.N; i++ {
		// code
	}
	b.ReportAllocs() // go test -benchmem
}

// b.Loop() 写法
func BenchmarkFoo(b *testing.B) {
	for b.Loop() {
		// code
	}
	b.ReportAllocs() // go test -benchmem
}
