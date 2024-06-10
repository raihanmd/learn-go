package utils

import (
	"fmt"
	"golang-basic/errors"
)

func HandleError() {
	err := recover()
	if err != nil {
		switch errorType := err.(type) {
		case *errors.ValidationError:
			fmt.Println("Validation Error:", errorType.Error())
		default:
			fmt.Println("Unknown Error:", err)
		}
	} else {
		fmt.Println("* ibarat sumber, & ibarat penunjuk")
	}
}
