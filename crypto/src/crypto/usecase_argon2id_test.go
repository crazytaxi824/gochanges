// mix of Argon2id + AES-GCM + HMAC-SHA256 encryption test
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

type aes struct {
	Algorithm string `json:"algo"`
	CipherHex string `json:"cipher"`
}

type record struct {
	Argon2id *crypto.Argon2Params `json:"argon2id"`
	AES      aes                  `json:"aes"`
}

func (r record) String() string {
	je, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(je)
}

func Argon2AESEncrypt(password, plaintext []byte, params *crypto.Argon2Params, algo string) (*record, error) {
	key, p, err := crypto.Argon2id(password, params)
	if err != nil {
		return nil, err
	}

	// AES encrypt
	var cipherBytes []byte
	switch algo {
	case "GCM":
		cipherBytes, err = crypto.AESGCMEncrypt(plaintext, key)
	case "CTR":
		cipherBytes, err = crypto.AESCTREncrypt(plaintext, key)
	case "CBC":
		cipherBytes, err = crypto.AESCBCEncrypt(plaintext, key)
	default:
		return nil, fmt.Errorf("algo error: %s", algo)
	}

	if err != nil {
		return nil, err
	}

	r := record{
		Argon2id: p,
		AES: aes{
			Algorithm: algo,
			CipherHex: hex.EncodeToString(cipherBytes),
		},
	}

	return &r, nil
}

func Argon2AESDecrypt(password []byte, rec *record) ([]byte, error) {
	key, _, err := crypto.Argon2id(password, rec.Argon2id)
	if err != nil {
		return nil, err
	}

	// 密文
	cipher, _ := hex.DecodeString(rec.AES.CipherHex)

	var plaintext []byte
	switch rec.AES.Algorithm {
	case "GCM":
		plaintext, err = crypto.AESGCMDecrypt(cipher, key)
	case "CTR":
		plaintext, err = crypto.AESCTRDecrypt(cipher, key)
	case "CBC":
		plaintext, err = crypto.AESCBCDecrypt(cipher, key)
	default:
		return nil, fmt.Errorf("algo error: %s", rec.AES.Algorithm)
	}

	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func TestArgon2AESEncrypt(t *testing.T) {
	password := []byte("password")
	plaintext := []byte("this is a AES test!!!")

	params := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    16,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	algo := "CTR"

	r, err := Argon2AESEncrypt(password, plaintext, params, algo)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(r)
}

func TestArgon2AESDecrypt(t *testing.T) {
	password := []byte("password")
	je := `{"argon2id":{"version":19,"memory":262144,"time":16,"threads":4,"key_len":32,"salt":"75db9c7f0cb06c3962153c0abe0748d3"},"aes":{"algo":"CTR","cipher":"00e12fa382a04cf4f034087b710d80f963e6690586d8f809674a42676fdec1cde99651e8f7"}}`

	var rec record
	err := json.Unmarshal([]byte(je), &rec)
	if err != nil {
		t.Error(err)
		return
	}

	plaintext, err := Argon2AESDecrypt(password, &rec)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(plaintext))
}
