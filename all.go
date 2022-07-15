package creek

// The All function determines whether all elements of the stream satisfy the passed condition.
func (s Stream[T]) All(expression func(T) bool) bool {
	return all(expression, s.Array)
}

// The All function determines whether all elements of the stream satisfy the passed condition.
func (s StructStream[T]) All(expression func(T) bool) bool {
	return all(expression, s.Array)
}

// The All function determines whether all elements of the stream satisfy the passed condition.
func (s MapStream[T, V]) All(expression func(KeyValuePair[T, V]) bool) bool {
	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			return false
		}
	}

	return true
}

// Function to reduce duplicated code.
func all[T interface{}](expression func(item T) bool, array []T) bool {
	for i := 0; i < len(array); i++ {
		if !expression(array[i]) {
			return false
		}
	}

	return true
}
