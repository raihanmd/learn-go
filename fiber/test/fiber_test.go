package test

import (
	"go-fiber/app"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = app.NewRouter()

func TestFiberApp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := router.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello, World!", string(bytes))
}
