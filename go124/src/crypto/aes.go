package crypto

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
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

func AESEncryptFile(plainFile, cipherFile *os.File, key []byte) (iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv, err = RandomBytes(block.BlockSize())
	if err != nil {
		return nil, err
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)

	freader := bufio.NewReader(plainFile)
	buffer := make([]byte, block.BlockSize())
	dst := make([]byte, block.BlockSize())

	var padded bool
	for {
		var n int
		n, err = freader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) && !padded {
				// 文件正好是 16 的倍数. 添加 []byte{16, 16 ... 16} pad
				buffer = pkcs7Pad([]byte{}, block.BlockSize())
				padded = true
			} else if errors.Is(err, io.EOF) && padded {
				// OK, EOF
				return iv, nil
			} else {
				return nil, err
			}
		}

		// if n > 0 && n < 16: end of file padding, mark padding
		if n > 0 && n < 16 {
			// end of file
			if padded {
				return nil, errors.New("encrypt Failed: padded more than once")
			}
			buffer = pkcs7Pad(buffer[:n], block.BlockSize())
			padded = true
		}

		// 加密数据
		encrypter.CryptBlocks(dst, buffer)

		// 分段写入文件
		_, err = cipherFile.Write(dst)
		if err != nil {
			return nil, err
		}
	}
}

func AESDecryptFile(plainFile, cipherFile *os.File, key, iv []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	decrypter := cipher.NewCBCDecrypter(block, iv)

	freader := bufio.NewReader(cipherFile)
	buffer := make([]byte, block.BlockSize())
	dst := make([]byte, block.BlockSize())

	var lastline bool
	for {
		var n int
		n, err = freader.Read(buffer)
		if err != nil && errors.Is(err, io.EOF) {
			// OK, EOF
			return nil
		} else if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		// 读取的数据必须是 16 的倍数.
		if n != block.BlockSize() {
			return errors.New("buffer read size is not 16 bytes")
		}

		// peak for EOF
		_, err = freader.Peek(1)
		if err != nil && errors.Is(err, io.EOF) {
			// next read is EOF, unpad current data.
			lastline = true
		} else if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		// 解密数据
		decrypter.CryptBlocks(dst, buffer[:n])
		if lastline {
			// unpad plain data
			dst, err = pkcs7Unpad(dst, block.BlockSize())
			if err != nil {
				return err
			}
		}

		// 写入文件
		_, err = plainFile.Write(dst)
		if err != nil {
			return err
		}
	}
}
