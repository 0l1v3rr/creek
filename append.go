package creek

// The Append function adds an element to the stream.
func (s Stream[T]) Append(item T) Stream[T] {
	s.Array = append(s.Array, item)

	return Stream[T]{
		Array: s.Array,
	}
}

// The AppendIf function adds an element to the stream if the second parameter is true.
func (s Stream[T]) AppendIf(item T, c bool) Stream[T] {
	if c {
		s.Array = append(s.Array, item)
	}

	return Stream[T]{
		Array: s.Array,
	}
}

// The AppendAt function inserts the specified element at the specified position in the stream.
func (s Stream[T]) AppendAt(index int, item T) Stream[T] {
	len := len(s.Array)
	if index >= len {
		return s
	}

	result := []T{}

	for i := 0; i < len; i++ {
		if i == index {
			result = append(result, item)
		}

		result = append(result, s.Array[i])
	}

	return Stream[T]{
		Array: result,
	}
}
