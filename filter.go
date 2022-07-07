package creek

// The Filter function leaves only those elements in the array
// that make the specified condition true.
func (s Stream[T]) Filter(expression func(item T) bool) Stream[T] {
	result := []T{}
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			result = append(result, s.Array[i])
		}
	}

	return Stream[T]{
		Array: result,
	}
}
