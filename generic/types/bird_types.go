package types

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
