package list

func CheckFor[T comparable](item T, from []T) bool {
	for _, element := range from {
		if element == item {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](item T, from []T) int {
	index := -1

	for i := 0; i < len(from); i++ {
		if from[i] == item {
			return i
		}
	}

	return index;
}
