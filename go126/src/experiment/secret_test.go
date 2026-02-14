// 执行后立即擦除栈与寄存器，防止私钥等敏感数据残留，增强前向保密性
// NOTE: 当前仅支持 Linux amd64/arm64

package experiment_test

import (
	"testing"

	"runtime/secret"
)

// `GOEXPERIMENT=runtimesecret go test -count=1 -v -run ^TestSecret$ local/src/experiment`
func TestSecret(t *testing.T) {
	secret.Do(func() {
		t.Log("omg")
	})
}
