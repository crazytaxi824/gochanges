package crypto_test

import (
	"bytes"
	"crypto/rand"
	"io"
	mrand "math/rand/v2"
	"testing"
)

func randBytes(size int) ([]byte, error) {
	randByts := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, randByts)
	if err != nil {
		return nil, err
	}

	for i := range randByts {
		randByts[i] = randByts[i] % 128 // 0~127
	}
	return randByts, nil
}

func Encode(data []byte, blockSize int) ([]byte, error) {
	res := make([]byte, 0)
	for _, v := range data {
		prefix, err := randBytes(blockSize - 1)
		if err != nil {
			return nil, err
		}

		p := append(prefix, v)
		res = append(res, p...)
	}

	suffix, err := randBytes(mrand.IntN(blockSize - 2))
	if err != nil {
		return nil, err
	}

	res = append(res, suffix...)
	return res, nil
}

func Decode(enc []byte, blockSize int) []byte {
	res := make([]byte, 0)
	i := 0
	for i+blockSize <= len(enc) {
		res = append(res, enc[i+blockSize-1]) // 每段最后一个字节
		i += blockSize
	}
	return res
}

func TestByteFormat(t *testing.T) {
	fSize := 6

	src := []byte("SNUDfdshhkhuHFUDIHSFUI")
	enc, err := Encode(src, fSize)
	if err != nil {
		t.Error(err)
		return
	}

	// t.Log(string(enc))

	dec := Decode(enc, fSize)
	if !bytes.Equal(src, dec) {
		t.Error("Not Equal")
	}
}
