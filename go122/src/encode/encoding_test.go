package encode_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestHexEncode(t *testing.T) {
	d := make([]byte, 8)
	t.Log(hex.Encode(d, []byte{200, 100, 0}))
	t.Log(string(d))
}

func TestHexAppendEncode(t *testing.T) {
	d := []byte("abc")
	s := []byte("xyz")
	r := hex.AppendEncode(d, s) // dst 的部分不变, 后面 hex.Encode() 后 append

	t.Log(d, s)
	t.Log(string(r))
}

func TestBase64Encode(t *testing.T) {
	d := make([]byte, 8)
	base64.StdEncoding.Encode(d, []byte{200, 100, 0})
	t.Log(string(d))
}

func TestBase64AppendEncode(t *testing.T) {
	d := []byte("abc")
	s := []byte("xyz")
	r := base64.StdEncoding.AppendEncode(d, s) // dst 的部分不变, 后面 hex.Encode() 后 append

	t.Log(d, s)
	t.Log(string(r))
}
