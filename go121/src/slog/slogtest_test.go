package slog_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"testing"
	"testing/slogtest"
)

// slogtest 测试 handler 是否设置正确
func TestSlogtestHandler(t *testing.T) {
	// 生成一个测试用的 JsonHandler
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)

	// 测试 handler 格式是否 json 格式.
	testResults := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			fmt.Println(string(line))
			if len(line) == 0 {
				continue
			}

			// 测试 json.Unmarshal() 是否正确.
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}

	// TestHandler() 会自动发送测试数据, 然后和 testResults() 返回数据做对比, 来判断 handler 是否正确.
	err := slogtest.TestHandler(handler, testResults)
	if err != nil {
		t.Fatal(err)
	}
}
