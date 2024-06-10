package utils

import "fmt"

type Person struct {
	FirstName, LastName string `required:"true"`
}

func SayHello(firstName, lastName string) {
	fmt.Printf("Hello \"%v %v\"! \n", firstName, lastName)
}

func PrintStruct(person *Person) {
	fmt.Printf("The val of struct without field name: %v\n", *person)
	fmt.Printf("The val of struct with field name: %+v\n", *person)
	fmt.Printf("The val of struct with field name and type of struct: %#v\n", *person)
	fmt.Printf("The type struct: %T\n", *person)
}
