package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/argon2"
)

// 参数结构体（可按需调整）
type Argon2Params struct {
	Memory  uint32 `json:"memory"`  // KB
	Time    uint32 `json:"time"`    // iterations, 迭代次数
	Threads uint8  `json:"threads"` // 并行计算数量
	KeyLen  uint32 `json:"key_len"` // key 长度, 如果要配合 AES 则需要使用 16|24|32
}

type Argon2Env struct {
	Version int `json:"version"` // argon2 version
	Argon2Params
	SaltHex string `json:"salt"` // 盐, 不要重复使用, 每次生成新的盐
}

// 盐的长度 16 足够
const saltLen = 16

// Argon2id 使用 Argon2id 生成哈希并以可移植字符串返回
func Argon2id(password []byte, params *Argon2Params, salt []byte) (key []byte, env *Argon2Env, err error) {
	if salt == nil {
		// 生成随机 salt
		salt = make([]byte, saltLen)
		if _, err = io.ReadFull(rand.Reader, salt); err != nil {
			return nil, nil, err
		}
	}

	// 派生密钥
	key = argon2.IDKey(password, salt, params.Time, params.Memory, params.Threads, params.KeyLen)

	// 返回结果
	return key, &Argon2Env{
		Version: argon2.Version,
		SaltHex: hex.EncodeToString(salt),
		Argon2Params: Argon2Params{
			Memory:  params.Memory,
			Time:    params.Time,
			Threads: params.Threads,
			KeyLen:  params.KeyLen,
		},
	}, nil
}
