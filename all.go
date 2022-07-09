package creek

// The All function determines whether all elements of the stream satisfy the passed condition.
func (s Stream[T]) All(expression func(item T) bool) bool {
	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			return false
		}
	}

	return true
}
