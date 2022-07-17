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

// The Map function creates a new stream populated with the results
// of calling the provided function on every element.
func (s MapStream[T, V]) Map(expression func(item KeyValuePair[T, V]) KeyValuePair[T, V]) MapStream[T, V] {
	result := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		result = append(result, expression(s.Array[i]))
	}

	return MapStream[T, V]{
		Array: result,
	}
}

func mapp[T interface{}](expression func(item T) T, arr []T) []T {
	result := []T{}

	for i := 0; i < len(arr); i++ {
		result = append(result, expression(arr[i]))
	}

	return result
}
