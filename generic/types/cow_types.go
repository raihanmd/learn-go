package types

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Cow[T string | constraints.Float] struct {
	Name T
}

func (c *Cow[_]) Eat() {
	fmt.Printf("%v eating grass\n", c.Name)
}

func (c *Cow[_]) Move() {
	fmt.Printf("%v is walking\n", c.Name)
}

func (c *Cow[_]) Speak() {
	fmt.Printf("%v mooing\n", c.Name)
}

func (c *Cow[_]) Hit(name string) {
	fmt.Printf("%v got hit by %v\n", c.Name, name)
}

func (c *Cow[T]) ChangeName(name T) {
	fmt.Printf("%v changed name to %v\n", c.Name, name)
	c.Name = name
	c.Speak()
}
