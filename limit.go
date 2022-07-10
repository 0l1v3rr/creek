package creek

// The Limit function constrains the number of elements returned by the stream.
func (s Stream[T]) Limit(amount int) Stream[T] {
	return Stream[T]{
		Array: limit(s.Array, amount),
	}
}

// The Limit function constrains the number of elements returned by the stream.
func (s StructStream[T]) Limit(amount int) StructStream[T] {
	return StructStream[T]{
		Array: limit(s.Array, amount),
	}
}

// Function to reduce duplicated code.
func limit[T interface{}](array []T, amount int) []T {
	if amount < 1 {
		return []T{}
	}

	if amount >= len(array) {
		return array
	}

	res := []T{}
	for i := 0; i < amount; i++ {
		res = append(res, array[i])
	}

	return res
}
