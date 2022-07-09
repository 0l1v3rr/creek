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