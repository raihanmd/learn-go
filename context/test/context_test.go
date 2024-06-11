package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	// ? TODO is the same as background
	todo := context.TODO()
	fmt.Println(ctx)
	fmt.Println(todo)
	assert.Equal(t, 1, 1)
}
