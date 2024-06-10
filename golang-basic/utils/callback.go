package utils

import "fmt"

func Callback(name string, filter func(string) bool) {
	if filter(name) {
		fmt.Println("You are blocked " + name)
	} else {
		fmt.Println("Hello, " + name)
	}

}
