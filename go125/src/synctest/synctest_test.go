package synctest

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// 依赖真实时间, 需要真实等待
	time.Sleep(3 * time.Second)
	if ctx.Err() != context.DeadlineExceeded {
		t.Fatalf("expected deadline exceeded")
	}
}

func TestSyncTestContextTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		// 时间是虚拟的，不会真的等
		time.Sleep(3 * time.Second)
		synctest.Wait()

		if ctx.Err() != context.DeadlineExceeded {
			t.Fatalf("expected deadline exceeded")
		}
	})
}

func TestSyncTestGoroutine(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan int, 1)
		go func() {
			ch <- 42
		}()

		synctest.Wait() // 等待 goroutine 进入稳定状态

		select {
		case v := <-ch:
			if v != 42 {
				t.Fatal("wrong value")
			}
		case <-time.After(3 * time.Second): // 时间是虚拟的，不会真的等
			t.Fatal("timeout")
		}
	})
}

func TestSyncTestTimer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		// 因为 synctest 控制时间，Sleep 不会真的等
		time.Sleep(3 * time.Second)
		synctest.Wait()

		select {
		case <-ticker.C:
			t.Log("OK") // 确认 tick 触发
		default:
			t.Fatal("expected tick")
		}
	})
}
