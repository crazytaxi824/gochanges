package slog_test

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
)

// 简单用法
func TestSlog(*testing.T) {
	// 简单用法
	slog.Info("hello")

	// 或者 ...args 必须是 kv pairs
	// slog.Info("hello", "count", 3, "bar") // Error
	slog.Info("hello", "count", 3, "bar", "foo")

	// 推荐用法, Int(), String() 中都必须是 KV pair.
	slog.Info("hello", slog.Int("count", 3), slog.String("foo", "bar"))

	// group args
	slog.Info("hello", slog.Group("contacts", slog.String("email", "abc@gg.com"), slog.String("mobile", "0284934343")))

	// 以下三个函数运行层面完全相同.
	slog.Info("msg")
	slog.Default().Info("msg")
	slog.Default().Log(context.Background(), slog.LevelInfo, "msg")
}

// slog.Level type 是 int
// LevelDebug Level = -4
// LevelInfo  Level = 0
// LevelWarn  Level = 4
// LevelError Level = 8
func TestSlogLevel(*testing.T) {
	l := slog.LevelInfo
	fmt.Println(l == l.Level())                                   // true
	fmt.Println(slog.LevelInfo == slog.LevelInfo.Level().Level()) // true
}
