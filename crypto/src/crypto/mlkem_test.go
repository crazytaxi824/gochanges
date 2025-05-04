// mlkem 是一种后量子密码学算法，主要用于密钥封装和交换。类似 Diffie-Hellman.
// Diffie-Hellman 和 mlkem 都不能防止 MIM 攻击.

package crypto_test

import (
	"crypto/mlkem"
	"encoding/hex"
	"testing"
)

func TestMlkem(t *testing.T) {
	// Alice
	k, err := mlkem.GenerateKey768()
	if err != nil {
		t.Log(err)
		return
	}

	// give to Bob
	ekbyte := k.EncapsulationKey().Bytes() // 如果中间人攻击获取 ekbyte 则不安全.

	// Bob
	ek, err := mlkem.NewEncapsulationKey768(ekbyte)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(ekbyte) == hex.EncodeToString(ek.Bytes())) // true

	// return cipher to Alice
	skb, cipher := ek.Encapsulate()
	t.Log(hex.EncodeToString(skb))

	// Alice
	ska, err := k.Decapsulate(cipher)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(ska))
}
