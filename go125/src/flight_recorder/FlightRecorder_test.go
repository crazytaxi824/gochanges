// https://go.dev/blog/execution-traces-2024
package flightrecorder

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
	"testing"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func TestFlightRecorder(t *testing.T) {
	// 新建记录
	cfg := trace.FlightRecorderConfig{
		MinAge:   5 * time.Second,
		MaxBytes: 5 << 20, // 5MB, MaxBytes is the priority
	}

	fr := trace.NewFlightRecorder(cfg)
	if er := fr.Start(); er != nil {
		t.Error(er)
		return
	}
	defer fr.Stop()

	var once sync.Once

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 使用新的 synctest 避免 wait 时间
		if !fr.Enabled() {
			t.Error("FlightRecorder is not Enabled")
			return
		}

		start := time.Now()

		// do heavy work()
		fib := fibonacci(40)
		t.Log(r.RemoteAddr)
		w.Write([]byte("fibonacci: " + strconv.Itoa(fib)))

		// trigger
		if time.Since(start) > 300*time.Millisecond {
			// 记录一次数据
			once.Do(func() {
				// Grab the snapshot.
				var b bytes.Buffer
				_, err := fr.WriteTo(&b)
				if err != nil {
					t.Error(err)
					return
				}
				// Write it to a file.
				// 使用 go tool trace ... 查看数据
				if err := os.WriteFile("trace.out", b.Bytes(), 0o755); err != nil {
					log.Print(err)
					return
				}
			})
		}
	})

	// 开启服务
	t.Error(http.ListenAndServe(":18080", mux))
}
