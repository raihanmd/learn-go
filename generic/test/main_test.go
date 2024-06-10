package test

import (
	"generic/types"
	"generic/utils"
	"testing"

	_ "github.com/gookit/goutil/testutil/assert"
)

func TestSum(t *testing.T) {
	utils.Equal(t, 10, utils.Sum[complex128](3.14, 2.22, 1.11, 3.53))

	result := utils.Sum(1, 2, 3, 4, 5)
	utils.Equal(t, 15, result)
}

func TestForceMove(t *testing.T) {
	cow := &types.Cow{}
	utils.Equal(t, &types.Cow{}, cow)
	utils.ForceMove(cow)
}
