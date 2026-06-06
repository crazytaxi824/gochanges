package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func AESCTREncrypt(plaintext, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// NOTE: CTR 模式的 IV 必须是 16 字节
	// VVI: IV 必须每次不同
	iv, err := RandomBytes(block.BlockSize())
	if err != nil {
		return nil, err
	}

	// CTR 模式的 IV 必须是 16 字节
	stream := cipher.NewCTR(block, iv)

	// CTR 模式的 IV 必须是 16 字节
	out := make([]byte, len(plaintext))
	stream.XORKeyStream(out, plaintext)

	// concat iv+out
	ciphertext = append(ciphertext, iv...)
	ciphertext = append(ciphertext, out...)
	return ciphertext, nil
}

func AESCTRDecrypt(ciphertext, key []byte) (plaintext []byte, err error) {
	iv, encryptedData := ciphertext[:16], ciphertext[16:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// CTR 模式的 IV 必须是 16 字节
	stream := cipher.NewCTR(block, iv)

	// CTR 模式的 IV 必须是 16 字节
	out := make([]byte, len(encryptedData))
	stream.XORKeyStream(out, encryptedData)
	return out, nil
}
