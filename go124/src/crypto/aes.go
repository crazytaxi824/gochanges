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

// AES 密钥长度只可以是16、24或32字节，分别对应AES-128、AES-192和AES-256, (1byte = 8bit, 32byte = 256bit)
// AES 块大小（block size）始终为 16 字节（128 位），无论是 AES-128、AES-192 还是 AES-256
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
	// data size 必须是 16 的倍数.
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid padded data length")
	}

	// 0 < pad size <= 16
	padLen := int(data[len(data)-1])
	if padLen <= 0 || padLen > blockSize {
		return nil, errors.New("invalid padding size")
	}

	// 验证 padding format. []byte{x,x,x,x,x,x,x,x,x,x,x,5,5,5,5,5}
	padding := data[len(data)-padLen:]
	if !bytes.Equal(padding, bytes.Repeat([]byte{byte(padLen)}, padLen)) {
		return nil, errors.New("invalid PKCS7 padding")
	}
	return data[:len(data)-padLen], nil
}

// 加密
func AESEncrypt(plaintext, key []byte) (ciphertext, iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	// AES: the only valid block-size is 128 bits(16 bytes)
	// fmt.Println(block.BlockSize())

	// 创建一个随机初始化向量 (IV), 长度必须为 BlockSize (16bytes) 固定值
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
