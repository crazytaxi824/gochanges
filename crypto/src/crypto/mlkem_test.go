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
	secretKey, err := mlkem.GenerateKey768()
	if err != nil {
		t.Log(err)
		return
	}

	// 生成 public key, 仅凭公钥 public key, 无法恢复出 shared key.
	pkByte := secretKey.EncapsulationKey().Bytes()
	t.Log(len(pkByte)) // 1184

	// give public key to Bob

	// Bob
	// get public key
	pk, err := mlkem.NewEncapsulationKey768(pkByte)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(pkByte) == hex.EncodeToString(pk.Bytes())) // true

	// 通过 public key 生成 cipher, 和 shared key.
	// 就算 cipher 也泄漏了, 同样无法恢复出 shared key.
	sharedKeyB, cipher := pk.Encapsulate()
	// t.Log(hex.EncodeToString(cipher))
	t.Log(len(cipher)) // 1088

	// return cipher to Alice

	// Alice
	// 通过 cipher 生成 shared key.
	sharedKeyA, err := secretKey.Decapsulate(cipher)
	if err != nil {
		t.Log(err)
		return
	}

	// shared key 一样, 使用对称加密.
	t.Log(hex.EncodeToString(sharedKeyB))
	t.Log(hex.EncodeToString(sharedKeyA))
}
