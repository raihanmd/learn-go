package test

import (
	"testing"

	"github.com/raihanmd/dependency_injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}
