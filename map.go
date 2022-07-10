package creek

// The Map function creates a new stream populated with the results
// of calling the provided function on every element.
func (s Stream[T]) Map(expression func(item T) T) Stream[T] {
	return Stream[T]{
		Array: mapp(expression, s.Array),
	}
}

// The Map function creates a new stream populated with the results
// of calling the provided function on every element.
func (s StructStream[T]) Map(expression func(item T) T) StructStream[T] {
	return StructStream[T]{
		Array: mapp(expression, s.Array),
	}
}

func mapp[T interface{}](expression func(item T) T, arr []T) []T {
	result := []T{}

	for i := 0; i < len(arr); i++ {
		result = append(result, expression(arr[i]))
	}

	return result
}
