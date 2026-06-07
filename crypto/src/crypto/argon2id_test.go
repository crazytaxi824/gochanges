package crypto_test

import (
	"bytes"
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
	_, env, err := crypto.Argon2id(password, defaultParams, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", env)
}

func TestArgon2KeyVarify(t *testing.T) {
	password := []byte("password")

	params := &crypto.Argon2Params{
		Memory:  256 * 1024, // KB
		Time:    12,         // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	// 生成哈希
	key1, env, err := crypto.Argon2id(password, params, nil)
	if err != nil {
		t.Error(err)
		return
	}

	salt, err := hex.DecodeString(env.SaltHex)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成哈希
	key2, _, err := crypto.Argon2id(password, &env.Argon2Params, salt)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(key1, key2) {
		t.Error("keys are not match")
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
		_, _, err := crypto.Argon2id(password, params, nil)
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.ReportAllocs() // go test -benchmem
}
