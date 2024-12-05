package lib

func IndexInSlice[T any](index int, slice []T) bool {
	return index >= 0 && index < len(slice)
}
