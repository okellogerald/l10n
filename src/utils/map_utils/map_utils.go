package maputils

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
