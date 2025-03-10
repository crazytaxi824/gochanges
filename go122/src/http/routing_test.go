// Routing Enhancements for Go 1.22
// https://go.dev/blog/routing-enhancements
package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "/index.html"`)
	})
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "GET /static/"`)
	})
	mux.HandleFunc("GET /foo/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "GET /foo/"`)
	})
	mux.HandleFunc("POST /foo/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "POST /foo/"`)
	})
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", func(w http.ResponseWriter, req *http.Request) {
		bucket := req.PathValue("bucket")
		objectname := req.PathValue("objectname")
		fmt.Fprintln(w, `match /b/{bucket}/o/{objectname...}, `+"bucket="+bucket+", objectname="+objectname)
	})

	return mux
}

func TestReq(t *testing.T) {
	mux := router()

	req := httptest.NewRequest("GET", "/index.html", http.NoBody)
	resp := httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // match "/index.html"

	req = httptest.NewRequest("POST", "/index.html", http.NoBody)
	resp = httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // match "/index.html"

	req = httptest.NewRequest("GET", "/static/", http.NoBody)
	resp = httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // match "GET /static/"

	req = httptest.NewRequest("POST", "/static/", http.NoBody)
	resp = httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // Method Not Allowed

	req = httptest.NewRequest("GET", "/foo/bar", http.NoBody)
	resp = httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // GET /foo/

	req = httptest.NewRequest("POST", "/foo/bar", http.NoBody)
	resp = httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // POST /foo/
}

func TestReq2(t *testing.T) {
	mux := router()

	req := httptest.NewRequest("GET", "/b/foo/o/bar/abc/def", http.NoBody)
	resp := httptest.NewRecorder()
	mux.ServeHTTP(resp, req)
	t.Log(resp.Body.String()) // match /b/{bucket}/o/{objectname...}, bucket=foo, objectname=bar/abc/def
}
