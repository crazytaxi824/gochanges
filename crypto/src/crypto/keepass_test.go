package crypto

import (
	"crypto/sha256"
	"testing"
)

func TestParseKeyFile(t *testing.T) {
	keyFilepath := "~/test.txt"

	hash, err := ParseKeyFile(keyFilepath)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(hash)
}

func TestCompositeKey2(t *testing.T) {
	password := []byte("123")
	keyFileHash := sha256.Sum256([]byte("456")) // 模拟一个 keyfile 的 sha256 hash

	compositeKey := GenCompositeKey(password, keyFileHash[:])
	if len(compositeKey) != sha256.Size {
		t.Error("compositeKey hash error")
	}

	t.Log(compositeKey)
}

func TestCompositeKey(t *testing.T) {
	password := []byte("123")
	keyFilepath := "~/test.txt"

	keyFileHash, err := ParseKeyFile(keyFilepath)
	if err != nil {
		t.Error(err)
		return
	}

	compositeKey := GenCompositeKey(password, keyFileHash)
	if len(compositeKey) != sha256.Size {
		t.Error("compositeKey hash error")
	}

	t.Log(compositeKey)
}
