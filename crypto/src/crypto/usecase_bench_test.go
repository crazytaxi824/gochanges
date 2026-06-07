package crypto_test

import (
	"bytes"
	"math/rand/v2"
	"testing"

	"local/src/crypto"
)

type ITesting interface {
	Log(...any)
	Logf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
}

func usecase(t ITesting, password, plain_orig []byte, algo string) {
	params := &crypto.Argon2Params{
		Memory:  64, // KB, 减少计算量
		Time:    16, // iterations
		Threads: 4,
		KeyLen:  32, // key 长度, 如果要配合 AES 则需要使用 16|24|32
	}

	rec, err := Argon2AESEncrypt(password, plain_orig, params, algo)
	if err != nil {
		t.Error(err)
		return
	}

	// 解密
	plain_dec, err := Argon2AESDecrypt(password, rec)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(plain_orig, plain_dec) {
		t.Error("decrypt failed")
		return
	}
}

func BenchmarkUsercase(b *testing.B) {
	var (
		GCM int
		CTR int
		CBC int
	)

	for b.Loop() {
		password, _ := crypto.RandomBytes(rand.IntN(12) + 6)
		plain_byts, _ := crypto.RandomBytes(rand.IntN(107) + 1)

		algoInt := rand.UintN(100) % 3
		var algoStr string
		switch algoInt {
		case 0:
			GCM += 1
			algoStr = "GCM"
		case 1:
			CTR += 1
			algoStr = "CTR"
		case 2:
			CBC += 1
			algoStr = "CBC"
		default:
			b.Errorf("algo error: %d", algoInt)
			return
		}

		usecase(b, password, plain_byts, algoStr)
	}
	b.ReportAllocs() // go test -benchmem
	b.Logf("GCM: %d, CTR: %d, CBC: %d", GCM, CTR, CBC)
}
