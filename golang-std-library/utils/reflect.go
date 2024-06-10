package utils

import "reflect"

func GetTypeName(obj any) any {
	return reflect.TypeOf(obj)
}

func IsValid(obj any) (result bool) {
	result = true
	var t reflect.Type = reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		var f reflect.StructField = t.Field(i)
		if f.Tag.Get("required") == "true" {
			result = reflect.ValueOf(obj).Field(i).Interface() != ""
			if !result {
				break
			}
		}
	}
	return result
}
