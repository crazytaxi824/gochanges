// 主要用于密码派生, 例如: signal protocol 使用 HKDF 通过持久密钥材料不断生成新的会话密钥，确保每次对话加密的安全性。
// 通过已有的密钥(高强度,长)派生出可选长度的密钥, 主要用于短期 session 使用.

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
	keyLen := 64

	b, err := hkdf.Key(sha3.New512, secret, salt, keyInfo, keyLen)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(b))

	b, _ = hkdf.Key(sha3.New512, secret, salt, keyInfo, 32)
	t.Log(hex.EncodeToString(b)) // 生成的 key 前半段是一样的
}

// NOTE: 一般情况下直接使用 Key()
func TestExtract(t *testing.T) {
	secret := []byte("secret")
	salt := []byte("random_salt") // 推荐长度至少为16字节

	prk, err := hkdf.Extract(sha3.New512, secret, salt)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(prk))
	t.Log(len(prk))

	okm, err := hkdf.Expand(sha3.New256, prk, "info", 48)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(okm))
	t.Log(len(okm))
}
