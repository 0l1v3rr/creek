package creek

// The Map function creates a new stream populated with the results
// of calling the provided function on every element.
func (s Stream[T]) Map(expression func(item T) T) Stream[T] {
	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		result = append(result, expression(s.Array[i]))
	}

	return Stream[T]{
		Array: result,
	}
}
