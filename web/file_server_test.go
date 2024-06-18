package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileServer(t *testing.T) {
	fileServer := http.FileServer(http.Dir("./resources"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	req := httptest.NewRequest(http.MethodGet, "/static/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Equal(t, "Hello World\r\n", string(body))
}

//go:embed resources
var resources embed.FS

func TestEmbed(t *testing.T) {
	dir, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	req := httptest.NewRequest(http.MethodGet, "/static/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Equal(t, "Hello World\r\n", string(body))
}

//go:embed resources/dashboard.html
var dashboardHTML string

func TestServeFile(t *testing.T) {
	ServeFile := func(w http.ResponseWriter, r *http.Request) {
		if cookie, _ := r.Cookie("session"); cookie.Value == "admin" {
			// ? Can with go embed
			fmt.Fprint(w, dashboardHTML)
		} else {
			http.ServeFile(w, r, "./resources/index.html")
		}
	}

	t.Run("Admin test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		req.AddCookie(&http.Cookie{Name: "session", Value: "admin"})
		ServeFile(rec, req)

		body, _ := io.ReadAll(rec.Result().Body)

		assert.Equal(t, "Hello admin\r\n", string(body))
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("User test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		req.AddCookie(&http.Cookie{Name: "session", Value: "user"})
		ServeFile(rec, req)

		body, _ := io.ReadAll(rec.Result().Body)

		assert.Equal(t, "Hello World\r\n", string(body))
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})
}
