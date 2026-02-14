package crypto_test

import (
	"crypto/rsa"
	"testing"
)

// 包括: dsa, ecdh, ecdsa, ed25519, mlken, tls, x509 ...
// 使用 `GODEBUG=cryptocustomrand=1 go test -v local/src/crypto` 兼容模式运行时会报错, io.Reader 不能为 nil.
func TestRSA(t *testing.T) {
	pk, err := rsa.GenerateKey(nil, 1024)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(pk.Public())
}
