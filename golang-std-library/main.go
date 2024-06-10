package main

import (
	"errors"
	"fmt"
	"golang-std-library/sorter"
	"golang-std-library/utils"
	"reflect"
	"slices"
)

var (
	ErrValidation = errors.New("validation error")
	ErrUnknown    = errors.New("unknown error")
)

func main() {
	var unsorted = []int{10, 64, -12, 23, 547, 12, 4, 53, 2}
	unsorted = slices.Insert(unsorted, len(unsorted), 100)
	fmt.Println(unsorted)
	fmt.Println(*sorter.BubbleSort(&unsorted))

	utils.SayHello("John", "Doe")

	Adit := &utils.Person{FirstName: "Adit", LastName: "Firman"}

	utils.PrintStruct(Adit)

	ErrCheck()

	fmt.Println(utils.GetArgs())

	fmt.Println(utils.GetHostname())

	utils.Flag()

	fmt.Println("========== Reverse String ==========")
	fmt.Println(utils.StrReverseLower("EVRESID"), utils.StrReverseUpper("dlrow olleh"))
	fmt.Print("========== Reverse String ==========\n\n")

	fmt.Println("========== Strconv ==========")
	fmt.Println(utils.StrToInt("123"), utils.StrToInt64("123"), utils.StrToFloat32("123.45"))
	fmt.Print("========== Strconv ==========\n\n")

	fmt.Println("========== Math ==========")
	fmt.Println(utils.IsPrime(40), utils.IsPrime(41), utils.IsPrime(29311))
	fmt.Print("========== Math ==========\n\n")

	fmt.Println("========== Time ==========")
	fmt.Println(utils.GetTimeNow().Day(), utils.GetTimeNow().Month(), utils.GetTimeNow().Year())
	fmt.Print("========== Time ==========\n\n")

	fmt.Println("========== Reflect ==========")
	fmt.Println(utils.GetTypeName(*Adit).(reflect.Type).Field(0).Name, utils.GetTypeName(*Adit).(reflect.Type).Field(0).Type, "\n", utils.IsValid(*Adit), utils.IsValid(utils.Person{FirstName: "Adit"}))
	fmt.Print("========== Reflect ==========\n\n")

	fmt.Println("========== CSV Reader ==========")
	utils.CsvReader("1,2,3,4,5,6,7,8,9,10\n11,12,13,14,15,16,17,18,19,20")
	fmt.Print("========== CSV Reader ==========\n\n")

	fmt.Println("========== Slices ==========")
	fmt.Println(utils.SortWithSlices([]int{10, 64, -12, 23, 547, 12, 4, 53, 2}))
	fmt.Println(slices.Min([]int{10, 64, -12, 23, 547, 12, 4, 53, 2}))
	fmt.Print("========== Slices ==========\n\n")

	fmt.Println("========== File ==========")
	fmt.Println("Writing to ./out/file.txt add content: Hello World! 123123")
	utils.CreateFile("out/file.txt", "Hello World!\n123123\n")
	utils.AddContentFile("out/file.txt", "Baru\n")
	utils.CreateFile("json/data.json", "{\"hello\": \"world\"}")
	resultFile, _ := utils.ReadFile("./out/file.txt")
	resultJson, _ := utils.ReadFile("./json/data.json")
	fmt.Println(resultFile)
	fmt.Println(resultJson)
	fmt.Println("========== File ==========")
}

func GetById(id string) error {
	if id == "" {
		return ErrValidation
	}

	if id != "Adit" {
		return ErrUnknown
	}

	return nil
}

func ErrCheck() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, ErrUnknown) {
			fmt.Println("Unknown Error")
		}
	}
}
