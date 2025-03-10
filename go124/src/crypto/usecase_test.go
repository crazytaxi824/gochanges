// mix of PBKDF2 with AES-CBC encryption test

package crypto_test

import (
	"crypto/pbkdf2"
	"crypto/sha3"
	"encoding/hex"
	"fmt"
	"testing"

	"local/src/crypto"
)

func TestMixEncrypt(t *testing.T) {
	password := "password"
	salt, _ := crypto.RandomBytes(32) // 推荐长度 > 16-bytes, 保证唯一性
	iter := 1000000
	keyLen := 32

	// pbkdf2 gen key
	key, err := pbkdf2.Key(sha3.New512, password, salt, iter, keyLen)
	if err != nil {
		t.Error(err)
		return
	}

	// AES encrypt
	plaintext := []byte("this is a AES test!!!")
	ciphertext, iv, err := crypto.AESEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("--- PBKDF2 keygen success --- ")
	fmt.Println("algo: PBKDF2-SHA3-512")               // 记录算法
	fmt.Printf("iter: %d\n", iter)                     // 记录 iteration
	fmt.Printf("salt: %s\n", hex.EncodeToString(salt)) // 记录 salt

	fmt.Println("--- AES encrypt success ---")
	fmt.Printf("algo: AES-%d-%s\n", len(key)*8, "CBC-PKCS7Padding") // 记录算法
	fmt.Printf("iv: %s\n", hex.EncodeToString(iv))                  // 记录 init vector
	fmt.Printf("cipher: %s\n", hex.EncodeToString(ciphertext))      // 记录 cipher text
}

func TestMixDecrypt(t *testing.T) {
	password := "password"
	salt, _ := hex.DecodeString("salt")
	iter := 1000000
	keyLen := 32

	// pbkdf2 gen key
	key, err := pbkdf2.Key(sha3.New512, password, salt, iter, keyLen)
	if err != nil {
		t.Error(err)
		return
	}

	// AES decrypt
	cipher, _ := hex.DecodeString("cipher text")
	iv, _ := hex.DecodeString("init vector")
	plaintext, err := crypto.AESDecrypt(cipher, iv, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(plaintext))
}
