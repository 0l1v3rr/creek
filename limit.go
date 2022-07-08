package creek

// The Limit function constrains the number of elements returned by the stream.
func (s Stream[T]) Limit(amount int) Stream[T] {
	if amount < 1 {
		return Stream[T]{
			Array: []T{},
		}
	}

	if amount >= len(s.Array) {
		return s
	}

	res := []T{}
	for i := 0; i < amount; i++ {
		res = append(res, s.Array[i])
	}

	return Stream[T]{
		Array: res,
	}
}
