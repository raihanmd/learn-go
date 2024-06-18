package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {

	if cookie, err := r.Cookie("session"); err == nil {
		fmt.Fprint(w, cookie.Value)
		return
	}

	cookie := new(http.Cookie)

	cookie.Name = "session"
	cookie.Value = "12345"
	cookie.Path = "/"

	http.SetCookie(w, cookie)
}

func TestCookie(t *testing.T) {
	t.Run("Without Cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		CookieHandler(rec, req)

		cookies := rec.Result().Cookies()

		assert.Equal(t, "12345", cookies[0].Value)
	})

	t.Run("With Cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "session", Value: "new_cookie", Path: "/"})
		rec := httptest.NewRecorder()

		CookieHandler(rec, req)

		res := rec.Result()

		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, "new_cookie", string(body))
	})

}
