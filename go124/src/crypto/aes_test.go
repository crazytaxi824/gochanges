// AES 密钥长度可以是16、24或32字节，分别对应AES-128、AES-192和AES-256。
// 1byte = 8bit, 32byte = 256bit

package crypto

import (
	"bytes"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestAESexample(t *testing.T) {
	// AES 密钥长度只可以是16、24或32字节，分别对应AES-128、AES-192和AES-256。
	// 1byte = 8bit, 32byte = 256bit
	key, _ := RandomBytes(32)

	// 要加密的数据
	plaintext := []byte("this is a AES test!!!")

	// 加密
	ciphertext, iv, err := AESEncrypt(plaintext, key)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(iv))
	t.Log(hex.EncodeToString(ciphertext))

	// 解密
	plaintext2, err := AESDecrypt(ciphertext, iv, key)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(plaintext2))

	// 验证
	t.Log(bytes.Equal(plaintext2, plaintext))
}

func TestAESFile(t *testing.T) {
	home, _ := os.LookupEnv("HOME")
	f, err := os.Open(filepath.Join(home, "Desktop", "testfile"))
	if err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		t.Log(err)
		return
	}

	key, _ := RandomBytes(32)
	cipherBytes, iv, err := AESEncrypt(content, key)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(hex.EncodeToString(iv))
	t.Log(len(cipherBytes))
	t.Log(hex.EncodeToString(cipherBytes[:64]))

	plainBytes, err := AESDecrypt(cipherBytes, iv, key)
	if err != nil {
		t.Log(err)
		return
	}

	nf, err := os.OpenFile(filepath.Join(home, "Desktop", "n.kdbx"), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		t.Log(err)
		return
	}
	defer nf.Close()

	_, err = nf.Write(plainBytes)
	if err != nil {
		t.Log(err)
		return
	}
}
