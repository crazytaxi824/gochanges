package crypto_test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"testing"
)

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	// fmt.Println(padding, len(padText), padText)  // padding == len(padText)
	return append(data, padText...)
}

func TestAESRangeEncrypt(t *testing.T) {
	block, _ := aes.NewCipher([]byte("passwordpasswordpasswordpassword"))
	iv := bytes.Repeat([]byte("a"), block.BlockSize())
	paddedText := pkcs7Pad([]byte("abc"), block.BlockSize())

	// 整体一次性加密
	ciphertext := make([]byte, len(paddedText))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext, paddedText)
	t.Log(hex.EncodeToString(ciphertext))

	// 按照 blocksize 分段加密
	block2, _ := aes.NewCipher([]byte("passwordpasswordpasswordpassword"))
	ciphertext2 := make([]byte, len(paddedText))
	encrypter2 := cipher.NewCBCEncrypter(block2, iv)
	for i := range len(paddedText) / block.BlockSize() {
		encrypter2.CryptBlocks(ciphertext2, paddedText[i*block.BlockSize():(i+1)*block.BlockSize()])
	}
	t.Log(hex.EncodeToString(ciphertext2))
}
