package creek

// The RemoveWhere function removes all the entries that satisfy the provided condition.
func (s Stream[T]) RemoveWhere(expression func(item T) bool) Stream[T] {
	result := []T{}
	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			result = append(result, s.Array[i])
		}
	}

	return Stream[T]{
		Array: result,
	}
}
