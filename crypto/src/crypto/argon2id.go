package crypto

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

// 参数结构体（可按需调整）
type Argon2Params struct {
	Version int    `json:"version"`  // argon2 version
	Memory  uint32 `json:"memory"`   // KB
	Time    uint32 `json:"time"`     // iterations, 迭代次数
	Threads uint8  `json:"threads"`  // 并行计算数量
	KeyLen  uint32 `json:"key_len"`  // key 长度, 如果要配合 AES 则需要使用 16|24|32
	SaltHex string `json:"salt_hex"` // 盐, 不要重复使用, 每次生成新的盐
}

// 盐的长度 16 足够
const saltLen = 16

// Argon2id 使用 Argon2id 生成哈希并以可移植字符串返回
func Argon2id(password []byte, params *Argon2Params) (key []byte, env *Argon2Params, err error) {
	var salt []byte
	if params.SaltHex == "" {
		// 生成随机 salt
		salt = make([]byte, saltLen)
		if _, err = rand.Read(salt); err != nil {
			return nil, nil, err
		}
	} else {
		if salt, err = hex.DecodeString(params.SaltHex); err != nil {
			return nil, nil, err
		}
	}

	// 派生密钥
	key = argon2.IDKey(password, salt, params.Time, params.Memory, params.Threads, params.KeyLen)

	// 返回结果
	return key, &Argon2Params{
		Version: argon2.Version,
		Memory:  params.Memory,
		Time:    params.Time,
		Threads: params.Threads,
		KeyLen:  params.KeyLen,
		SaltHex: hex.EncodeToString(salt),
	}, nil
}
