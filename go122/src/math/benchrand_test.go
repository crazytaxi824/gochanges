package math_test

import (
	crand "crypto/rand"
	"math/big"
	"math/rand/v2"
	"testing"
	"time"
)

// 伪随机-1，获取一次seed，不停生成随机数，如果seed被破解，会有重大安全隐患。
func BenchmarkPseudoRandomInt(b *testing.B) {
	r := rand.New(rand.NewPCG(uint64(time.Now().Unix()), 1))
	for i := 0; i < b.N; i++ {
		_ = r.IntN(10000)
	}
	b.ReportAllocs()
}

// 伪随机-2，每次重新获取 time 作为seed
func BenchmarkPseudoRandomInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := rand.New(rand.NewPCG(uint64(time.Now().Unix()), 1))
		_ = r.IntN(10000)
	}
	b.ReportAllocs()
}

// 真随机测试
func BenchmarkTrueRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := crand.Int(crand.Reader, big.NewInt(10000))
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.ReportAllocs()
}
