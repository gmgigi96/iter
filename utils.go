package iter

// zero return the zero value for the given type.
func zero[T any]() T {
	var t T
	return t
}
