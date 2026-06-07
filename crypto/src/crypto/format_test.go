package crypto_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	mrand "math/rand/v2"
	"slices"
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

func Encode(data []byte, fmtSize int) ([]byte, error) {
	if fmtSize < 3 {
		return nil, fmt.Errorf("size is too small: %d", fmtSize)
	}

	res := make([]byte, 0)
	for _, v := range data {
		prefix, err := randBytes(fmtSize - 1)
		if err != nil {
			return nil, err
		}

		res = slices.Concat(res, prefix, []byte{v})
	}

	suffix, err := randBytes(mrand.IntN(fmtSize - 2))
	if err != nil {
		return nil, err
	}

	res = append(res, suffix...)
	return res, nil
}

func Decode(enc []byte, fmtSize int) ([]byte, error) {
	if fmtSize < 3 {
		return nil, fmt.Errorf("size is too small: %d", fmtSize)
	}

	res := make([]byte, 0)
	i := 0
	for i+fmtSize <= len(enc) {
		res = append(res, enc[i+fmtSize-1]) // 每段最后一个字节
		i += fmtSize
	}
	return res, nil
}

func TestByteFormat(t *testing.T) {
	fSize := 6

	src := []byte("ABCDEFGHIJK")
	enc, err := Encode(src, fSize)
	if err != nil {
		t.Error(err)
		return
	}

	// t.Log(string(enc))

	dec, err := Decode(enc, fSize)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(src, dec) {
		t.Error("Not Equal")
	}
}
