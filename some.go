package creek

// The Some function determines whether any of the elements of the stream satisfy the passed condition.
func (s Stream[T]) Some(expression func(item T) bool) bool {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return true
		}
	}

	return false
}

// The Some function determines whether any of the elements of the stream satisfy the passed condition.
func (s StructStream[T]) Some(expression func(item T) bool) bool {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return true
		}
	}

	return false
}

// The Some function determines whether any of the elements of the stream satisfy the passed condition.
func (s MapStream[T, V]) Some(expression func(item KeyValuePair[T, V]) bool) bool {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return true
		}
	}

	return false
}
