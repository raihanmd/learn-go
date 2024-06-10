package utils

import "strconv"

func StrToInt(str string) int {
	var result, _ = strconv.Atoi(str)
	return result
}

func StrToInt64(str string) int64 {
	var result, _ = strconv.ParseInt(str, 10, 64)
	return result
}

func StrToFloat32(str string) float64 {
	var result, _ = strconv.ParseFloat(str, 32)
	return result
}
