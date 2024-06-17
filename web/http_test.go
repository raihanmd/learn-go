package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	if firstName == "" && lastName == "" {
		fmt.Fprint(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s %s!", firstName, lastName)
	}
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	QueryHandler(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Hello, World!", string(res))
}

func TestQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?first_name=Adit&last_name=Firmansyah", nil)
	rec := httptest.NewRecorder()

	QueryHandler(rec, req)

	res, _ := io.ReadAll(rec.Result().Body)

	assert.Equal(t, "Hello, Adit Firmansyah!", string(res))
}
