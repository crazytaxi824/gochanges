package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func RandomBytes(l int) ([]byte, error) {
	r := make([]byte, l)
	_, err := io.ReadFull(rand.Reader, r)
	return r, err
}

// 如果 data 正好是 blockSize 的倍数, 则 PKCS#7 会在 data 最后 pad 16 个 []byte(16)
// pkcs7Pad 对明文数据进行PKCS#7填充
// 如需要填充 5 个数, 则为 [5]byte{5,5,5,5,5}
// 如需要填充 3 个数, 则为 [3]byte{3,3,3}
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	// fmt.Println(padding, len(padText), padText)  // padding == len(padText)
	return append(data, padText...)
}

// pkcs7Unpad
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	padding := int(data[length-1])
	if padding <= 0 || padding > blockSize {
		return nil, errors.New("invalid padding")
	}
	return data[:length-padding], nil
}

// 加密
func AESEncrypt(plaintext, key []byte) (ciphertext, iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	// AES: the only valid block-size is 128 bits(16 bytes)
	// fmt.Println(block.BlockSize())

	// 创建一个随机初始化向量 (IV), 长度必须为 BlockSize - 16 bytes 固定值
	// NOTE: IV 必须每次不同.
	iv, err = RandomBytes(block.BlockSize())
	if err != nil {
		return nil, nil, err
	}

	// data after padding
	paddedText := pkcs7Pad(plaintext, block.BlockSize())

	// 使用 AES-CBC 模式进行加密
	ciphertext = make([]byte, len(paddedText))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext, paddedText)

	return ciphertext, iv, nil
}

// 必须要知道 IV 才能正确解密.
func AESDecrypt(ciphertext, iv, key []byte) (plaintext []byte, err error) {
	// 创建一个 AES 块密码
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 解密
	paddedText := make([]byte, len(ciphertext))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(paddedText, ciphertext)

	// unpad data
	return pkcs7Unpad(paddedText, block.BlockSize())
}
