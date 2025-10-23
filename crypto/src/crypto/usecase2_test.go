// mix of Argon2id + AES-CBC + HMAC-SHA256 encryption test
// Argon2id: 派生密钥
// AES-CBC:  加密数据
// HAMC: 验证数据完成性

package crypto_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"local/src/crypto"
)

type aes_gcm struct {
	Cipher string `json:"cipher"`
}

type record struct {
	Argon2id *crypto.Argon2Params `json:"argon2id"`
	AES_GCM  aes_gcm              `json:"aes_gcm"`
}

func (r record) String() string {
	je, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(je)
}

func TestMixEncrypt2(t *testing.T) {
	password := []byte("password")
	plaintext := []byte("this is a AES test!!!")

	params := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    16,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}
	key, p, err := crypto.Argon2id(password, params)
	if err != nil {
		t.Error(err)
		return
	}

	// AES encrypt
	cipherBytes, err := crypto.AESGCMEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}

	r := record{
		Argon2id: p,
		AES_GCM: aes_gcm{
			Cipher: hex.EncodeToString(cipherBytes),
		},
	}
	fmt.Println(r)
}

func TestMixDecrypt2(t *testing.T) {
	password := []byte("password")
	je := `{argon2id & cipher to decode}`

	var rec record
	err := json.Unmarshal([]byte(je), &rec)
	if err != nil {
		t.Error(err)
		return
	}

	key, _, err := crypto.Argon2id(password, rec.Argon2id)
	if err != nil {
		t.Error(err)
		return
	}

	// 密文
	cipher, _ := hex.DecodeString(rec.AES_GCM.Cipher)
	plaintext, err := crypto.AESGCMDecrypt(cipher, key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(plaintext))
}
