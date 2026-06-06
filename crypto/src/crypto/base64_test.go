package crypto_test

import (
	"encoding/base64"
	"testing"
)

func TestBase64(t *testing.T) {
	src := []byte{18, 19, 20, 21}
	t.Log(base64.StdEncoding.EncodeToString(src))
	t.Log(base64.URLEncoding.EncodeToString(src))    // 符号 + 和 / 会被替换为 - 和 _
	t.Log(base64.RawStdEncoding.EncodeToString(src)) // 不会在最后添加 = 作为填充

	encod := base64.RawStdEncoding.EncodeToString(src)

	dst, err := base64.RawStdEncoding.DecodeString(encod)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(dst))
}
