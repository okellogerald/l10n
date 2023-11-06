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

	return index
}

func IndexWhere[T comparable](from []T, predicate func(T) bool) int {
	index := -1 // Initialize index to -1

	for i := 0; i < len(from); i++ {
		if predicate(from[i]) { // If predicate function returns true
			return i // Return current index
		}
	}

	return index // Return index if not found
}

func Where[T comparable](from []T, predicate func(T) bool) []T {
	var list []T

	for i := 0; i < len(from); i++ {
		if predicate(from[i]) {
			list = append(list, from[i])
		}
	}

	return list
}

func Combine[T comparable](list1 []T, list2 []T) []T {
	result := make([]T, len(list1))
	copy(result, list1)

	result = append(result, list2...)
	return result
}
