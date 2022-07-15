package creek

// The ForEach method runs the specified method with every element in the Stream.
func (s Stream[T]) ForEach(expression func(item T)) {
	for i := 0; i < len(s.Array); i++ {
		expression(s.Array[i])
	}
}

// The ForEach method runs the specified method with every element in the Stream.
func (s StructStream[T]) ForEach(expression func(item T)) {
	for i := 0; i < len(s.Array); i++ {
		expression(s.Array[i])
	}
}

// The ForEach method runs the specified method with every element in the Stream.
func (s MapStream[T, V]) ForEach(expression func(KeyValuePair[T, V])) {
	for i := 0; i < len(s.Array); i++ {
		expression(s.Array[i])
	}
}
