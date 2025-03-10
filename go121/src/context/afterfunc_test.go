package context_test

import (
	"context"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	stop := context.AfterFunc(ctx, func() {
		t.Log("AfterFunc")
	})

	// stop context 执行 AfterFunc 内的函数. stop 之后再 cancel 不会触发 AfterFunc.
	stop()
	cancel()
	time.Sleep(1 * time.Second)
}

func TestAfterFunc2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	stop := context.AfterFunc(ctx, func() {
		// cancel() 的时候执行
		t.Log("AfterFunc")
	})

	cancel()
	stop()
	time.Sleep(1 * time.Second)
}
