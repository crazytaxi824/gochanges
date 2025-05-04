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
	// 生成 key
	password := "password"
	salt, _ := crypto.RandomBytes(32) // 推荐长度 > 16-bytes, 保证唯一性
	iter := 1000000                   // 迭代次数, 增加暴力破解难度
	keyLen := 32                      // 生成密钥长度

	// 明文, 需要加密
	plaintext := []byte("this is a AES test!!!")

	// pbkdf2 gen key
	key, err := pbkdf2.Key(sha3.New512, password, salt, iter, keyLen)
	if err != nil {
		t.Error(err)
		return
	}

	// AES encrypt
	ciphertext, err := crypto.AESGCMEncrypt(plaintext, key)
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
	fmt.Printf("cipher: %s\n", hex.EncodeToString(ciphertext))      // 记录 cipher text
}

func TestMixDecrypt(t *testing.T) {
	// 生成 key
	password := "password"
	salt, _ := hex.DecodeString("salt hex")
	iter := 1000000
	keyLen := 32

	// 密文
	cipher, _ := hex.DecodeString("cipher hex")

	// pbkdf2 gen key
	key, err := pbkdf2.Key(sha3.New512, password, salt, iter, keyLen)
	if err != nil {
		t.Error(err)
		return
	}

	// AES decrypt
	plaintext, err := crypto.AESGCMDecrypt(cipher, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(plaintext))
}
