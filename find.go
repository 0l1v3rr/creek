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

// The Find function searches for an element that matches the conditions passed
// and returns the first occurrence within the entire stream.
func (s StructStream[T]) Find(expression func(item T) bool) T {
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
func (s StructStream[T]) FindIndex(expression func(item T) bool) int {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}

// The FindLast function searches for an element that matches the conditions passed
// and returns the last occurrence within the entire stream.
func (s StructStream[T]) FindLast(expression func(item T) bool) T {
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
func (s StructStream[T]) FindLastIndex(expression func(item T) bool) int {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}

// The Find function searches for an element that matches the conditions passed
// and returns the first occurrence within the entire stream.
func (s MapStream[T, V]) Find(expression func(item KeyValuePair[T, V]) bool) KeyValuePair[T, V] {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return s.Array[i]
		}
	}

	return KeyValuePair[T, V]{}
}

// The FindIndex function searches for an element that matches the conditions passed
// and returns the index of the first occurrence within the entire stream.
func (s MapStream[T, V]) FindIndex(expression func(item KeyValuePair[T, V]) bool) int {
	for i := 0; i < len(s.Array); i++ {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}

// The FindLast function searches for an element that matches the conditions passed
// and returns the last occurrence within the entire stream.
func (s MapStream[T, V]) FindLast(expression func(item KeyValuePair[T, V]) bool) KeyValuePair[T, V] {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if expression(s.Array[i]) {
			return s.Array[i]
		}
	}

	return KeyValuePair[T, V]{}
}

// The FindLastIndex function searches for an element that matches the conditions passed
// and returns the index of the last occurrence within the entire stream.
func (s MapStream[T, V]) FindLastIndex(expression func(item KeyValuePair[T, V]) bool) int {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if expression(s.Array[i]) {
			return i
		}
	}

	return -1
}
