package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Golang")

	auth := r.Header.Get("Authorization")

	if auth != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	fmt.Fprintf(w, "Hello %s", auth)
}

func TestHttpHeader(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Authorization", "admin")

		rec := httptest.NewRecorder()

		AuthHandler(rec, req)

		res, _ := io.ReadAll(rec.Result().Body)

		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		assert.Equal(t, "Hello admin", string(res))
		assert.Equal(t, "Golang", rec.Result().Header.Get("x-powered-by"))
	})

	t.Run("Failed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Authorization", "user")

		rec := httptest.NewRecorder()

		AuthHandler(rec, req)

		res, _ := io.ReadAll(rec.Result().Body)

		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		assert.Equal(t, "Unauthorized", string(res))
		assert.Equal(t, "Golang", rec.Result().Header.Get("x-powered-by"))
	})

}
