package slog_test

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"testing"
	"time"
)

// 复杂设置
type MyLogger struct {
	id     int64
	logger *slog.Logger
}

func NewLogger(id int64, logger *slog.Logger) MyLogger {
	return MyLogger{
		id:     id,
		logger: logger,
	}
}

// 直接 copy logger.log() 源代码
func (l *MyLogger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.logger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, log])
	pc := pcs[0]

	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	r.AddAttrs(slog.Int64("id", l.id)) // 添加自己的属性

	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.logger.Handler().Handle(ctx, r)
}

func (l *MyLogger) Info(msg string, args ...any) {
	l.log(context.Background(), slog.LevelInfo, msg, args...)
}

func (l *MyLogger) Warn(msg string, args ...any) {
	l.log(context.Background(), slog.LevelWarn, msg, args...)
}

func (l *MyLogger) Error(msg string, args ...any) {
	l.log(context.Background(), slog.LevelError, msg, args...)
}

func (l *MyLogger) Debug(msg string, args ...any) {
	l.log(context.Background(), slog.LevelDebug, msg, args...)
}

func TestCustomRecord(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}))
	l := NewLogger(999, logger)

	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
}
