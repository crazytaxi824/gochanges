package crypto

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestKeepassKeyFile(t *testing.T) {
	// 生成 256bit (32byte) key 数据
	src := sha256.Sum256([]byte("123")) // VVI: 可以是任何数据

	// generate XML key file
	keyFile, err := FormatKeepassKeyFile(src[:])
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Println(string(keyFile))
}

func TestCompositeKey(t *testing.T) {
	password := []byte("123")
	keyFileHash := sha256.Sum256([]byte("456"))
	compositeKey := GenCompositeKey(password, keyFileHash[:])
	if len(compositeKey) != sha256.Size {
		t.Error("compositeKey hash error")
	}

	t.Log(compositeKey)
}
