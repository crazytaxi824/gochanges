// 主要用于密码派生, 例如: signal protocol 使用 HKDF 通过持久密钥材料不断生成新的会话密钥，确保每次对话加密的安全性。

package crypto_test

import (
	"crypto/hkdf"
	"crypto/sha3"
	"encoding/hex"
	"testing"
)

func TestHkdf(t *testing.T) {
	secret := []byte("secret")
	salt := []byte("random_salt") // 推荐长度至少为16字节
	keyInfo := "区分上下文, 用于区分不同用途的密钥，防止密钥重用"
	keyLen := 32

	b, err := hkdf.Key(sha3.New512, secret, salt, keyInfo, keyLen)
	if err != nil {
		t.Log(err)
		return
	}

	s := hex.EncodeToString(b)
	t.Log(s)
	t.Log(len(s))

	b, _ = hkdf.Key(sha3.New512, secret, salt, keyInfo, 20)
	t.Log(hex.EncodeToString(b)) // secret 的前半段 20 和 32 是一样的
}

// 一般情况下使用 Key()
func TestExtract(t *testing.T) {
	secret := []byte("secret")
	salt := []byte("random_salt") // 推荐长度至少为16字节

	b, err := hkdf.Extract(sha3.New512, secret, salt)
	if err != nil {
		t.Log(err)
		return
	}

	s := hex.EncodeToString(b)
	t.Log(s)
	t.Log(len(s))
}
