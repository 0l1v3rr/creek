package creek

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it throws a panic.
func (s Stream[T]) ElementAt(index int) T {
	return s.Array[index]
}

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it returns the elseValue, which is the second parameter.
func (s Stream[T]) ElementAtOrElse(index int, elseValue T) T {
	if index >= len(s.Array) {
		return elseValue
	}

	return s.Array[index]
}
