package test

import (
	"encoding/json"
	"go-fiber/app"
	"go-fiber/model/web"
	"go-fiber/model/web/response"
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

	var response web.WebSuccessResponse[response.User]

	bytes, _ := io.ReadAll(res.Body)
	json.Unmarshal(bytes, &response)

	assert.Equal(t, response.Code, 200)
	assert.Equal(t, response.Message, "OK")
	assert.Equal(t, response.Payload.ID, 1)
	assert.Equal(t, response.Payload.Name, "Raihanmd")
	assert.Equal(t, response.Payload.Email, "Raihanmd")
}

func TestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Add("Authorization", "admin")

	res, err := router.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}
