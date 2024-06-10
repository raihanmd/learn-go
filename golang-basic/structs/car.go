package structs

import "fmt"

type Car struct {
	Brand    string
	Speed    int
	Repaired bool
}

func (car *Car) Gas() {
	fmt.Println(car.Brand, "is running on gas, with speed", car.Speed)
}

func (car *Car) Repair() {
	car.Repaired = true
}
