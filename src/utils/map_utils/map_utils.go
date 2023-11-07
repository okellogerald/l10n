package maputils

import (
	"fmt"
	"reflect"
)

func Combine[T comparable, S comparable](map1, map2 map[T]S) map[T]S {
	combinedMap := make(map[T]S, len(map1)+len(map2))

	for key, value := range map1 {
		combinedMap[key] = value
	}

	for key, value := range map2 {
		combinedMap[key] = value
	}

	return combinedMap
}

func ConvertToContentMap(value any) (map[string]interface{}, bool) {
	println(fmt.Sprintf("%v", value))
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Map {
		myMap := make(map[string]interface{})
		iter := v.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			value := iter.Value().Interface()
			myMap[key] = value
		}
		return myMap, true
	} else {
		return nil, false
	}
}
