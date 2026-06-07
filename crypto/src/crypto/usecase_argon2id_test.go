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

func Argon2AESEncrypt(password, plainBytes []byte, params *crypto.Argon2Params, algo string, fmtSize int) (*record, error) {
	key, p, err := crypto.Argon2id(password, params)
	if err != nil {
		return nil, err
	}

	enc, err := Encode(plainBytes, fmtSize)
	if err != nil {
		return nil, err
	}

	// AES encrypt
	var cipherBytes []byte
	switch algo {
	case "GCM":
		cipherBytes, err = crypto.AESGCMEncrypt(enc, key)
	case "CTR":
		cipherBytes, err = crypto.AESCTREncrypt(enc, key)
	case "CBC":
		cipherBytes, err = crypto.AESCBCEncrypt(enc, key)
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

func Argon2AESDecrypt(password []byte, rec *record, fmtSize int) ([]byte, error) {
	key, _, err := crypto.Argon2id(password, rec.Argon2id)
	if err != nil {
		return nil, err
	}

	// 密文
	cipherBytes, err := hex.DecodeString(rec.AES.CipherHex)
	if err != nil {
		return nil, err
	}

	var plaintext []byte
	switch rec.AES.Algorithm {
	case "GCM":
		plaintext, err = crypto.AESGCMDecrypt(cipherBytes, key)
	case "CTR":
		plaintext, err = crypto.AESCTRDecrypt(cipherBytes, key)
	case "CBC":
		plaintext, err = crypto.AESCBCDecrypt(cipherBytes, key)
	default:
		return nil, fmt.Errorf("algo error: %s", rec.AES.Algorithm)
	}
	if err != nil {
		return nil, err
	}

	dec := Decode(plaintext, fmtSize)
	return dec, nil
}

func TestArgon2AESEncrypt(t *testing.T) {
	password := []byte("password")
	plaintext := []byte("this is a AES test!!!")
	algo := "CTR"
	fSize := 6

	params := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    16,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	r, err := Argon2AESEncrypt(password, plaintext, params, algo, fSize)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(r)
}

func TestArgon2AESDecrypt(t *testing.T) {
	password := []byte("password")
	je := `{"argon2id":{"version":19,"memory":262144,"time":16,"threads":4,"key_len":32,"salt":"527404a60f1c5580742d610262dec624"},"aes":{"algo":"CTR","cipher":"9e282c7c42bdab28182a7ee2961ef205b201e186c2dc75e22e2bc84b15fb3037595c926c090a4f046462fe6998c0c194b6f82a02c4438667b04492bbc4c60bba36eae5fee6ff2eac3f9c576118d43f5163363d05febf0893f9dc907ef6e14ef85cda2309e23868afec43e4fb1d42427315f5bd9557820a2571ec4bacc6951ef55da62d3c0001e4e0312c4a2a0341"}}`
	fSize := 6

	var rec record
	err := json.Unmarshal([]byte(je), &rec)
	if err != nil {
		t.Error(err)
		return
	}

	plaintext, err := Argon2AESDecrypt(password, &rec, fSize)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(plaintext))
}
