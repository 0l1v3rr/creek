package creek

// The Values function returns a new stream of the values of the map.
func (s MapStream[T, V]) Values() Stream[V] {
	res := []V{}

	for _, j := range s.Array {
		res = append(res, j.Value)
	}

	return Stream[V]{
		Array: res,
	}
}
