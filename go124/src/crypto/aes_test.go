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
	key, _ := RandomBytes(32)
	home, _ := os.LookupEnv("HOME")

	// source file
	sf, err := os.Open(filepath.Join(home, "Desktop", "testfile2"))
	if err != nil {
		t.Log(err)
		return
	}
	defer sf.Close()

	// cipher file
	cf, err := os.OpenFile(filepath.Join(home, "Desktop", "cipher"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		t.Log(err)
		return
	}
	defer cf.Close()

	// 加密 file
	iv, err := AESEncryptFile(sf, cf, key)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hex.EncodeToString(iv))
	cf.Seek(0, 0) // move offset back to beginning

	// plain file: decrypt from cipher file
	pf, err := os.OpenFile(filepath.Join(home, "Desktop", "plain"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		t.Log(err)
		return
	}
	defer pf.Close()

	err = AESDecryptFile(pf, cf, key, iv)
	if err != nil {
		t.Log(err)
		return
	}

	// compare two files
	sf.Seek(0, 0) // move offset back to beginning
	pf.Seek(0, 0)
	sfb, err := io.ReadAll(sf)
	if err != nil {
		t.Log(err)
		return
	}

	pfb, err := io.ReadAll(pf)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(len(sfb), len(pfb))
	t.Log(bytes.Equal(sfb, pfb))
}
