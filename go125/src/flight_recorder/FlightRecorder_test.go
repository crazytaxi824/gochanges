// https://go.dev/blog/execution-traces-2024
package flightrecorder

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestFlightRecorder(t *testing.T) {
	// 新建记录
	cfg := trace.FlightRecorderConfig{
		MinAge:   5 * time.Second,
		MaxBytes: 5 << 20, // 5MB
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
		synctest.Test(t, func(t *testing.T) {
			if !fr.Enabled() {
				t.Error("FlightRecorder is not Enabled")
				return
			}

			start := time.Now()

			// dowork()
			time.Sleep(1 * time.Second)
			synctest.Wait()
			t.Log(r.RemoteAddr)
			w.Write([]byte("hello world"))

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
	})

	// 开启服务
	t.Error(http.ListenAndServe(":18080", mux))
}
