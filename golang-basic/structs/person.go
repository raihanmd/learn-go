package structs

type Person struct {
	FirstName, Lastname string
	Age                 int
	Hobby               []string
}

func (person Person) GetName() string {
	return person.FirstName
}
