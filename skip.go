package creek

// The Skip function discards the first n elements of a stream, where n is the passed parameter.
func (s Stream[T]) Skip(amount int) Stream[T] {
	return Stream[T]{
		Array: skip(amount, s.Array),
	}
}

// The Skip function discards the first n elements of a stream, where n is the passed parameter.
func (s StructStream[T]) Skip(amount int) StructStream[T] {
	return StructStream[T]{
		Array: skip(amount, s.Array),
	}
}

func skip[T interface{}](amount int, arr []T) []T {
	if amount < 1 {
		return arr
	}

	len := len(arr)

	if amount >= len {
		return []T{}
	}

	var res []T

	for i := amount; i < len; i++ {
		res = append(res, arr[i])
	}

	return res
}
