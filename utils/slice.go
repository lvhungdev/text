package utils

func SliceInsertAt[T any](slice []T, index int, el T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}

	if len(slice) == 0 {
		return []T{el}
	}

	if index == len(slice) {
		return append(slice, el)
	}

	result := append(slice[:index+1], slice[index:]...)
	result[index] = el

	return result
}

func SliceRemoveAt[T any](slice []T, index int) []T {
	if len(slice) == 0 || index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
