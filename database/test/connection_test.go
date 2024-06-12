package test

import (
	"golang_database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	conn := golang_database.CreateConnection()
	defer conn.Close()
	assert.NotNil(t, conn)
}
