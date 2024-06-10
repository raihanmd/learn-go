package main

import (
	"fmt"
	"golang-basic/errors"
	"golang-basic/structs"
	"golang-basic/utils"
)

func main() {
	defer utils.HandleError()
	myFunc := func() {
		fmt.Println("Hello, World!")
	}

	myFunc()

	utils.Callback("John", func(name string) bool {
		if name == "John" {
			return true
		} else {
			return false
		}
	})

	Aditya := structs.Person{
		FirstName: "Adit",
		Lastname:  "Firmansyah",
		Age:       17,
		Hobby:     []string{"Coding", "Football"},
	}

	Lynx := structs.Person{
		FirstName: "Lynx",
		Lastname:  "Dev",
		Age:       20,
		Hobby:     []string{"Coding"},
	}

	Mobil := structs.Car{
		Brand: "Mobil",
		Speed: 10,
	}

	fmt.Println(Mobil)

	Yamaha := &Mobil

	fmt.Println(Yamaha)

	changeCarToIcikiwir(Yamaha)

	Yamaha.Repair()

	Mobil.Gas()

	fmt.Println("Mobil brand: ", Mobil.Brand)

	fmt.Println(Mobil)

	fmt.Println(Aditya)
	fmt.Println(Lynx)

	res := random()
	resultStr := res.(string)
	fmt.Println(resultStr)

	myError := errors.ValidationError{Message: "Idk why error"}

	fmt.Println(myError)

	SaveData("")
}

func random() interface{} {
	return "OK"
}

func changeCarToIcikiwir(car *structs.Car) {
	car.Brand = "Icikiwir"
}

func SaveData(data string) {
	if data == "" {
		panic(&errors.ValidationError{Message: "Data cannot be an empty string!"})
	}
	fmt.Println("Save Data Sukses")
}
