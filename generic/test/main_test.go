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

func TestStructGeneric(t *testing.T) {
	cow := &types.Cow[string]{Name: "Derandro"}
	cow01 := &types.Cow[float32]{Name: 0.1}
	bird := &types.Bird{}
	bird.Speak()
	cow.Eat()
	cow.Speak()
	cow.ChangeName("Chickenado")
	cow.Hit("Eko")
	utils.Equal(t, "Chickenado", cow.Name)
	utils.Equal(t, 0.1, cow01.Name)
	utils.ForceMove(cow)
	utils.ForceMove(cow01)
}

func TestTypeGeneric(t *testing.T) {
	fruits := types.Inventory[string]{}
	fruits.Add("Apple")
	fruits.Add("Orange")
	fruits.Print()

	primeNumbers := types.Inventory[int]{}
	primeNumbers.Add(2)
	primeNumbers.Add(3)
	primeNumbers.Print()

	utils.Equal(t, 2, len(fruits))
	utils.Equal(t, 2, primeNumbers[0])
}

// ? Generic can also used in interface
