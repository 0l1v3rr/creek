package creek

// The Push function adds the passed array to the end of the stream.
func (s Stream[T]) Push(arr []T) Stream[T] {
	s.Array = append(s.Array, arr...)
	return s
}

// The Push function adds the passed array to the end of the stream.
func (s StructStream[T]) Push(arr []T) StructStream[T] {
	s.Array = append(s.Array, arr...)
	return s
}

// The PushValues function adds the passed values to the end of the stream.
func (s Stream[T]) PushValues(values ...T) Stream[T] {
	s.Array = append(s.Array, values...)
	return s
}

// The PushValues function adds the passed values to the end of the stream.
func (s StructStream[T]) PushValues(values ...T) StructStream[T] {
	s.Array = append(s.Array, values...)
	return s
}
