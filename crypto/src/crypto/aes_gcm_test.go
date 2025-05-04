package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"testing"
)

func TestAESGCMEncrypt(t *testing.T) {
	key := bytes.Repeat([]byte("password"), 4)
	plaintext := []byte("this is a AES-256-GCM test!!!")
	ciphertext, err := AESGCMEncrypt(plaintext, key)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(hex.EncodeToString(ciphertext))
}

func TestAESGCMDecrypt(t *testing.T) {
	key := bytes.Repeat([]byte("password"), 4)
	ciphertext, _ := hex.DecodeString("ciphertext")

	plaintext, err := AESGCMDecrypt(ciphertext, key)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(string(plaintext))
}

func TestAESGCMEncrypt2(t *testing.T) {
	key := bytes.Repeat([]byte("password"), 4)
	plaintext := []byte("this is a AES-256-GCM test!!!")
	aad := bytes.Repeat([]byte("abcd"), 100)

	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)

	// 创建 random nonce(iv), nonce 一定不能重复, 否则会有严重安全隐患.
	nonce := make([]byte, aesGCM.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	// 加密并附加认证标签
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, aad)

	// 解密
	block2, _ := aes.NewCipher(key)
	aesGCM2, _ := cipher.NewGCM(block2)
	ns := aesGCM2.NonceSize()
	p, _ := aesGCM2.Open(nil, ciphertext[:ns], ciphertext[ns:], aad)

	t.Log(bytes.Equal(plaintext, p))
}
