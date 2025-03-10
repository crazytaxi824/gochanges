// pbkdf2 主要用于需要通过用户密码生成密钥的场景.
// - 可以随意选择密钥长度.
// - 主要用于防止彩虹表攻击, 或其他方式的暴力破解.
// 许多加密货币钱包使用 PBKDF2 来保护私钥或种子短语（seed phrase）
//
// PBKDF2中的 iter 迭代次数和 salt 可以公开，不会降低安全性。
// - Salt 的安全性不依赖于它的保密性，而是依赖于它的随机性和唯一性,
//   即使攻击者知道salt值，他们仍然需要为每个用户的唯一salt创建新的彩虹表.
// - iter 迭代次数的安全性来自于计算量，而非保密性.
// 所以 salt 越长, iter 越大越不容易被破解.

package crypto_test

import (
	"crypto/pbkdf2"
	"crypto/sha3"
	"encoding/hex"
	"testing"

	"local/src/crypto"
)

func Pbkdf2(password string, keyLen int) (key, salt []byte, err error) {
	// 可以公开, 迭代次数的安全性来自于计算量，而非保密性.
	iter := 10000

	// 可以公开, Salt 的安全性不依赖于它的保密性，而是依赖于它的随机性和唯一性.
	// 即使攻击者知道salt值，他们仍然需要为每个用户的唯一salt创建新的彩虹表.
	// 获取随机的 salt. 长度越长, 唯一性越高. salt 长度对计算量影响的不大.
	salt, _ = crypto.RandomBytes(32) // 推荐长度 >16, 保证唯一性

	// sha3-512 比 sha256(sha2-256) 慢 5~6 倍.
	key, err = pbkdf2.Key(sha3.New512, password, salt, iter, keyLen)
	return key, salt, err
}

// password 生成密钥
func TestPbkdf2(t *testing.T) {
	key, salt, err := Pbkdf2("password", 16)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("salt: %s", hex.EncodeToString(salt))
	t.Logf("key len: %d", len(key))
	t.Logf("key: %s", hex.EncodeToString(key))
}
