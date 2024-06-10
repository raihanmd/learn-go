package utils

import "strings"

func StrReverseLower(str string) (result string) {
	for _, val := range str {
		result = strings.ToLower(string(val)) + result
	}
	return result
}

func StrReverseUpper(str string) (result string) {
	for _, val := range str {
		result = strings.ToUpper(string(val)) + result
	}
	return result
}
