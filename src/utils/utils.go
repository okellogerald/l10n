package utils

import (
	"reflect"
	"strings"
)

func ConvertTo[T any](fromValue any, toValue any) (T, bool) {
	reflectValue := reflect.ValueOf(fromValue)
	convertedValue := reflectValue.Convert(reflect.TypeOf(toValue))
	result, ok := convertedValue.Interface().(T)
	return result, ok
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	firstLetter := strings.ToUpper(string(s[0]))
	return firstLetter + s[1:]
}