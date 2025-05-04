package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// VVI: 默认 NonceSize = 12
// nonce(iv), additionalData 都可以公开, 但是 nonce 一定不能重复, 否则会有严重安全隐患.
func AESGCMEncrypt(plaintext, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建GCM模式
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 创建 random nonce(iv), nonce 一定不能重复, 否则会有严重安全隐患.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 加密并附加认证标签
	ciphertext = aesGCM.Seal(nonce, nonce, plaintext, nil)
	// 注意：返回的ciphertext包含nonce在前面，然后是实际密文

	return ciphertext, nil
}

func AESGCMDecrypt(ciphertext, key []byte) (plaintext []byte, err error) {
	// 创建cipher块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建GCM模式
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 验证密文长度
	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("cipher text size error: less than nonce size")
	}

	// 提取nonce并解密
	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err = aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
