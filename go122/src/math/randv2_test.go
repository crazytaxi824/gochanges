package math_test

import (
	crand "crypto/rand"
	"math/big"
	"math/rand/v2"
	"testing"
)

// 伪随机 Pseudo-Random
func TestRandV2(t *testing.T) {
	t.Log(rand.IntN(100)) // new seed every time

	r := rand.New(rand.NewPCG(0, 1)) // fixed seed
	t.Log(r.IntN(100))               // 11
	t.Log(r.IntN(100))               // 10
	t.Log(r.IntN(100))               // 1
}

// 真随机 True-Random
func TestCryptoRand(t *testing.T) {
	n, err := crand.Int(crand.Reader, big.NewInt(100))
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(n.String())
}
