package test

import (
	"fmt"
	"testing"
	"unit_test/src/utils"

	"github.com/gookit/goutil/testutil/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")
	m.Run()
	fmt.Println("All tests passed")
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name string
		req  int
		exp  bool
	}{
		{"IsPrime", 13, true},
		{"IsPrime", 878656757, false},
		{"IsPrime", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, utils.IsPrime(tt.req))
		})
	}
}
