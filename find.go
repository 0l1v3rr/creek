package creek

// The Find function searches for an element that matches the conditions passed
// and returns the first occurrence within the entire stream.
func (s Stream[T]) Find(expression func(item T) bool) T {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return s.Array[i]
		}
	}

	var res T
	return res
}

// The FindIndex function searches for an element that matches the conditions passed
// and returns the index of the first occurrence within the entire stream.
func (s Stream[T]) FindIndex(expression func(item T) bool) int {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}

// The FindLast function searches for an element that matches the conditions passed
// and returns the last occurrence within the entire stream.
func (s Stream[T]) FindLast(expression func(item T) bool) T {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if expression(s.Array[i]) {
			return s.Array[i]
		}
	}

	var res T
	return res
}

// The FindLastIndex function searches for an element that matches the conditions passed
// and returns the index of the last occurrence within the entire stream.
func (s Stream[T]) FindLastIndex(expression func(item T) bool) int {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}