package utils

import "reflect"

func ConvertTo[T any](fromValue any, toValue any) (T, bool) {
	reflectValue := reflect.ValueOf(fromValue)
	convertedValue := reflectValue.Convert(reflect.TypeOf(toValue))
	result, ok := convertedValue.Interface().(T)
	return result, ok
}
