package slog_test

import (
	"log/slog"
	"os"
	"slices"
	"testing"
)

// simple custom logger
func TestCustomLogger(t *testing.T) {
	// h := slog.NewJSONHandler(os.Stdout, nil)  // 使用默认 opts
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,           // 记录调用该 logger 的 filepath, line, function.
		Level:     slog.LevelWarn, // >= 该 Level 时才会 log, 默认是 0, 即: LevelInfo. LevelDebug = -4

		// 每个 Attr 都会运行一次该函数.
		// built-in group "source" 不会出现在 group 中, 而是出现在 attr 中.
		ReplaceAttr: func(group []string, attr slog.Attr) slog.Attr {
			t.Log(group, attr)

			// 将默认的 time 格式替换为 time unix.
			if len(group) == 0 && attr.Key == slog.TimeKey {
				return slog.Int64(slog.TimeKey, attr.Value.Time().Unix())
			}

			// built-in group "source" 不会出现在 group 中, 而是出现在 attr 中.
			if attr.Key == slog.SourceKey {
				attr.Key = "S" // 修改 "source" 为 "S"
				return attr
			}

			// remove attr,
			// slog.Group() 定义的 group 会出现在这里.
			if slices.Contains(group, "g") && attr.Key == "k1" {
				return slog.Attr{}
			}

			return attr
		},
	})

	// 生成 custom logger
	logger := slog.New(handler)

	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")

	// k1=v1, k2=v2 在 [g] 的 group 中.
	// group 返回 []string{"g"}
	logger.Error("error", slog.Group("g", slog.String("k1", "v1"), slog.String("k2", "v2")))

	// k1=v1, k2=v2 在 [w1 w2] 的 group 中.
	// group 返回 []string{"w1","w2"}
	logger.WithGroup("w1").WithGroup("w2").Error("with group",
		slog.String("k1", "v1"), slog.String("k2", "v2"))

	// k1=v1, k2=v2 在 [w1 w2 g] 的 group 中.
	// group 返回 []string{"w1","w2","g"}
	logger.WithGroup("w1").WithGroup("w2").Error("with group",
		slog.Group("g", slog.String("k1", "v1"), slog.String("k2", "v2")))
}
