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

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it throws a panic.
func (s StructStream[T]) ElementAt(index int) T {
	return s.Array[index]
}

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it returns the elseValue, which is the second parameter.
func (s StructStream[T]) ElementAtOrElse(index int, elseValue T) T {
	if index >= len(s.Array) {
		return elseValue
	}

	return s.Array[index]
}

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it throws a panic.
func (s MapStream[T, V]) ElementAt(index int) KeyValuePair[T, V] {
	return s.Array[index]
}

// The ElementAt function is used to get an element from the stream at a particular index.
// If the element is not present, it returns the elseValue, which is the second parameter.
func (s MapStream[T, V]) ElementAtOrElse(index int, elseValue KeyValuePair[T, V]) KeyValuePair[T, V] {
	if index >= len(s.Array) {
		return elseValue
	}

	return s.Array[index]
}

// The First method returns the first element in the stream.
func (s Stream[T]) First() T {
	return s.Array[0]
}

// The First method returns the first element in the stream.
func (s StructStream[T]) First() T {
	return s.Array[0]
}

// The First method returns the first element in the stream.
func (s MapStream[T, V]) First() KeyValuePair[T, V] {
	return s.Array[0]
}

// The Last method returns the last element in the stream.
func (s Stream[T]) Last() T {
	return s.Array[len(s.Array)-1]
}

// The Last method returns the last element in the stream.
func (s StructStream[T]) Last() T {
	return s.Array[len(s.Array)-1]
}

// The Last method returns the last element in the stream.
func (s MapStream[T, V]) Last() KeyValuePair[T, V] {
	return s.Array[len(s.Array)-1]
}
