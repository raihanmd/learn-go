package test

import (
	"testing"
	"unit_test/src/utils"

	"github.com/gookit/goutil/testutil/assert"
)

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, utils.MergeSort([]int{8, 7, 6, 5, 4, 3, 2, 1}))
}
