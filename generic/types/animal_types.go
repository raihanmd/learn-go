package types

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c *Cow) Eat() {
	println("I'm eating grass")
}

func (c *Cow) Move() {
	println("I'm walking")
}

func (c *Cow) Speak() {
	println("I'm mooing")
}

type Bird struct{}

func (b *Bird) Eat() {
	println("I'm eating seeds")
}

func (b *Bird) Move() {
	println("I'm flying")
}

func (b *Bird) Speak() {
	println("I'm pecking")
}
