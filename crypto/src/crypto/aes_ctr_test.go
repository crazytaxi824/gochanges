package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

func TestAESCTR(t *testing.T) {
	key, _ := RandomBytes(32)
	data := []byte("this is a AES CTR test!!!")

	ciphertext, err := AESCTREncrypt(data, key)
	if err != nil {
		t.Error(err)
		return
	}

	p, err := AESCTRDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(p))
	t.Log(bytes.Equal(p, data))
}

// CTR 在任何情况下都会解密出一段数据, 不会报错
func TestAESCTRDecrypt(t *testing.T) {
	key, _ := RandomBytes(32)
	data := []byte("this is a AES CTR test!!!")

	block, err := aes.NewCipher(key)
	if err != nil {
		t.Error(err)
	}

	// NOTE: CTR 模式的 IV 必须是 16 字节
	// VVI: IV 必须每次不同
	iv, err := RandomBytes(block.BlockSize())
	if err != nil {
		t.Error(err)
	}

	// CTR 模式的 IV 必须是 16 字节
	stream := cipher.NewCTR(block, iv)

	// CTR 模式的 IV 必须是 16 字节
	out := make([]byte, len(data))
	stream.XORKeyStream(out, data)

	t.Log(string(out))
}
