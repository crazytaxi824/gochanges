package crossorigin

import (
	"net/http"
	"testing"
	"time"
)

// 需要构造一个 POST 请求
func TestCrossOrigin(t *testing.T) {
	// 防御 CSRF 攻击
	c := http.NewCrossOriginProtection()

	// 只有来自这些来源的请求才会被允许
	err := c.AddTrustedOrigin("https://trusted.com")
	if err != nil {
		t.Log(err)
		return
	}
	err = c.AddTrustedOrigin("https://trusted2.com")
	if err != nil {
		t.Log(err)
		return
	}

	// 添加不受保护的路径
	c.AddInsecureBypassPattern("/a/")

	// handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.Header.Get("Origin")) // 获取 Origin
		t.Log(r.RemoteAddr)
	})
	mux.HandleFunc("/a/b", func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.Header.Get("Origin")) // 获取 Origin
		t.Log(r.RemoteAddr)
	})

	svr := &http.Server{
		Addr:         ":18080",
		Handler:      c.Handler(mux),
		ReadTimeout:  5 * time.Second,   // 读取超时时间
		WriteTimeout: 10 * time.Second,  // 写入超时时间
		IdleTimeout:  120 * time.Second, // 空闲连接超时
	}

	if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		t.Error("Server error:", err)
	}
}
