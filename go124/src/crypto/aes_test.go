// AES 密钥长度可以是16、24或32字节，分别对应AES-128、AES-192和AES-256。
// 1byte = 8bit, 32byte = 256bit

package crypto_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"local/src/crypto"
)

func TestAESexample(t *testing.T) {
	// AES 密钥长度只可以是16、24或32字节，分别对应AES-128、AES-192和AES-256。
	// 1byte = 8bit, 32byte = 256bit
	key, _ := crypto.RandomBytes(32)

	// 要加密的数据
	plaintext := []byte("this is a AES test!!!")

	// 加密
	ciphertext, iv, err := crypto.AESEncrypt(plaintext, key)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(iv))
	t.Log(hex.EncodeToString(ciphertext))

	// 解密
	plaintext2, err := crypto.AESDecrypt(ciphertext, iv, key)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(plaintext2))

	// 验证
	t.Log(bytes.Equal(plaintext2, plaintext))
}
