package utils

import (
	"generic/types"
	"testing"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer | constraints.Float | constraints.Complex](n ...T) (result T) {
	for _, val := range n {
		result += val
	}
	return result
}

func Equal[T comparable](t *testing.T, expect, actual T) {
	if expect != actual {
		t.Errorf("expect: %v, actual: %v", expect, actual)
	}
}

func ForceMove[T types.Animal](animal T) {
	animal.Move()
}
