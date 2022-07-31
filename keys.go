package creek

// The Keys function returns a new stream of the keys of the map.
func (s MapStream[T, V]) Keys() Stream[T] {
	res := []T{}

	for _, j := range s.Array {
		res = append(res, j.Key)
	}

	return Stream[T]{
		Array: res,
	}
}
