package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestAESGCMEncrypt(t *testing.T) {
	key := bytes.Repeat([]byte("password"), 4)
	plaintext := []byte("this is a AES-256-GCM test!!!")
	cipher, err := AESGCMEncrypt(plaintext, key)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(hex.EncodeToString(cipher))
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
