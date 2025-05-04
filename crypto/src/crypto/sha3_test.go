// sha256 is sha2

package crypto_test

import (
	"crypto/sha256"
	"crypto/sha3"
	"encoding/hex"
	"testing"
)

func TestSha3(t *testing.T) {
	h := sha3.New256()
	_, err := h.Write([]byte("abc"))
	if err != nil {
		t.Log(err)
		return
	}

	b := h.Sum(nil)
	t.Log(len(b))

	r := hex.EncodeToString(b)
	t.Log(r)
	t.Log(len(r))
}

// sha256 is sha2 Not Sha3
// `$ sha256 -s abc`
func TestSha256(t *testing.T) {
	h := sha256.New()
	_, err := h.Write([]byte("abc"))
	if err != nil {
		t.Log(err)
		return
	}

	b := h.Sum(nil)
	t.Log(len(b))

	r := hex.EncodeToString(b)
	t.Log(r)
	t.Log(len(r))
}
