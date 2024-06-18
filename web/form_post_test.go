package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FormPostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := r.PostForm.Get("name")
	address := r.PostForm.Get("address")

	// Or use PostFormValue or FormValue without ParseForm first

	// name := r.PostFormValue("name")
	// address := r.PostFormValue("address")

	if name == "" || address == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func TestFormPost(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		reqBody := strings.NewReader("name=Adit&address=Jakarta")
		req := httptest.NewRequest(http.MethodPost, "/", reqBody)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		FormPostHandler(rec, req)
		res, _ := io.ReadAll(rec.Result().Body)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		assert.Equal(t, "Name = Adit\nAddress = Jakarta\n", string(res))
	})

	t.Run("Failed with empty body", func(t *testing.T) {
		reqBody := strings.NewReader("name=Adit")
		req := httptest.NewRequest(http.MethodPost, "/", reqBody)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		FormPostHandler(rec, req)
		res, _ := io.ReadAll(rec.Result().Body)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
		assert.Equal(t, "Bad Request\n", string(res))
	})

	t.Run("Failed with wrong method", func(t *testing.T) {
		reqBody := strings.NewReader("name=Adit&address=Jakarta")
		req := httptest.NewRequest(http.MethodGet, "/", reqBody)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		FormPostHandler(rec, req)
		res, _ := io.ReadAll(rec.Result().Body)
		assert.Equal(t, http.StatusMethodNotAllowed, rec.Result().StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(res))
	})

}
