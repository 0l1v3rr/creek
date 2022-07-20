package creek

// The Bind function binds the stream into the passed variable.
func (s Stream[T]) Bind(v *[]T) {
	*v = s.Collect()
}

// The Bind function binds the stream into the passed variable.
func (s StructStream[T]) Bind(v *[]T) {
	*v = s.Collect()
}

// The Bind function binds the stream into the passed variable.
func (s MapStream[T, V]) Bind(v *map[T]V) {
	*v = s.Collect()
}
