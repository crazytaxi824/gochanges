package crypto_test

import (
	"encoding/hex"
	"testing"

	"local/src/crypto"
)

var defaultParams = &crypto.Argon2Params{
	Memory:  64 * 1024, // 64 MB (单位为 KB)
	Time:    3,         // iterations
	Threads: 4,
	KeyLen:  32,
}

func TestArgon2(t *testing.T) {
	password := []byte("password")

	// 生成哈希
	_, params, err := crypto.Argon2id(password, defaultParams)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", params)
}

func TestArgon2KeyVarify(t *testing.T) {
	password := []byte("password")

	p := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    12,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	// 生成哈希
	key1, params, err := crypto.Argon2id(password, p)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成哈希
	key2, _, err := crypto.Argon2id(password, params)
	if err != nil {
		t.Error(err)
		return
	}

	if hex.EncodeToString(key1) != hex.EncodeToString(key2) {
		t.Error("keys are not match")
	} else {
		t.Log("keys are match")
	}
}

func BenchmarkArgon2(b *testing.B) {
	password := []byte("password")
	params := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    16,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	for b.Loop() {
		// code
		_, _, err := crypto.Argon2id(password, params)
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.ReportAllocs() // go test -benchmem
}
