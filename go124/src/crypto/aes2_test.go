package crypto

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

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

func TestEncryptFile(t *testing.T) {
	var err error
	home, _ := os.LookupEnv("HOME")
	sf, err := os.Open(filepath.Join(home, "Desktop", "testfile"))
	if err != nil {
		t.Log(err)
		return
	}
	defer sf.Close()

	cf, err := os.OpenFile(filepath.Join(home, "Desktop", "cipher"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		t.Log(err)
		return
	}
	defer cf.Close()

	// f := strings.NewReader(strings.Repeat("a", 16))

	block, _ := aes.NewCipher([]byte("passwordpasswordpasswordpassword"))
	iv := bytes.Repeat([]byte("a"), block.BlockSize())
	encrypter := cipher.NewCBCEncrypter(block, iv)

	fr := bufio.NewReader(sf)
	buffer := make([]byte, block.BlockSize())
	dst := make([]byte, block.BlockSize())

	var padded bool
	for {
		var n int
		n, err = fr.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) && !padded {
				// 文件正好是 16 的倍数.
				buffer = pkcs7Pad([]byte{}, block.BlockSize())
				padded = true
			} else if errors.Is(err, io.EOF) && padded {
				// OK, EOF
				return
			} else {
				t.Log(err)
				return
			}
		}

		// if n > 0 && n < 16: end of file padding, mark padding
		if n > 0 && n < 16 {
			// end of file
			if padded {
				t.Log("encrypt Failed: padded more than once")
				return
			}
			buffer = pkcs7Pad(buffer[:n], block.BlockSize())
			padded = true
		}

		// t.Log(padded, len(buffer), buffer)
		encrypter.CryptBlocks(dst, buffer)

		_, err = cf.Write(dst)
		if err != nil {
			t.Log(err)
			return
		}
	}
}

func TestDecryptFile(t *testing.T) {
	var err error
	home, _ := os.LookupEnv("HOME")
	cf, err := os.Open(filepath.Join(home, "Desktop", "cipher"))
	if err != nil {
		t.Log(err)
		return
	}
	defer cf.Close()

	pf, err := os.OpenFile(filepath.Join(home, "Desktop", "plain.kdbx"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		t.Log(err)
		return
	}
	defer pf.Close()

	block, _ := aes.NewCipher([]byte("passwordpasswordpasswordpassword"))
	iv := bytes.Repeat([]byte("a"), block.BlockSize())
	decrypter := cipher.NewCBCDecrypter(block, iv)

	fr := bufio.NewReader(cf)
	buffer := make([]byte, block.BlockSize())
	dst := make([]byte, block.BlockSize())

	var lastline bool
	for {
		var n int
		n, err = fr.Read(buffer)
		if err != nil {
			t.Log(err)
			return
		}

		if n != block.BlockSize() {
			t.Log("buffer is not 16 bytes")
			return
		}

		_, err = fr.Peek(1)
		if err != nil && errors.Is(err, io.EOF) {
			lastline = true
		} else if err != nil && !errors.Is(err, io.EOF) {
			t.Log(err)
			return
		}

		decrypter.CryptBlocks(dst, buffer[:n])
		if lastline {
			dst, err = pkcs7Unpad(dst, block.BlockSize())
			if err != nil {
				t.Log(err)
				return
			}
		}

		// if len(dst) > 0 {
		// 	t.Log(dst)
		// }

		_, err = pf.Write(dst)
		if err != nil {
			t.Log(err)
			return
		}
	}
}
