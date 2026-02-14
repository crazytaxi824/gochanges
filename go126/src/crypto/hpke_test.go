// Hybrid Public Key Encryption. support for post-quantum hybrid KEMs.

package crypto_test

import (
	"testing"

	"crypto/hpke"
)

func TestHPKE(t *testing.T) {
	// KEM:  (密钥封装)
	// KDF:  (密钥派生)
	// AEAD: (AES 加密)
	kem, kdf, aead := hpke.MLKEM768X25519(), hpke.HKDFSHA256(), hpke.AES256GCM()

	// Recipient side, 生成密钥对
	var (
		recipientPrivateKey hpke.PrivateKey
		publicKeyBytes      []byte
	)
	{
		k, err := kem.GenerateKey()
		if err != nil {
			panic(err)
		}
		recipientPrivateKey = k
		publicKeyBytes = k.PublicKey().Bytes()
	}

	// Sender side, 通过接受方的公钥加密数据
	var ciphertext []byte
	{
		publicKey, err := kem.NewPublicKey(publicKeyBytes)
		if err != nil {
			panic(err)
		}

		message := []byte("|-()-|")
		ct, err := hpke.Seal(publicKey, kdf, aead, []byte("example"), message)
		if err != nil {
			panic(err)
		}

		ciphertext = ct
	}

	// Recipient side, 解密
	{
		plaintext, err := hpke.Open(recipientPrivateKey, kdf, aead, []byte("example"), ciphertext)
		if err != nil {
			panic(err)
		}
		t.Logf("Decrypted message: %s\n", plaintext)
	}
}
