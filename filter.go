package creek

// The Filter function leaves only those elements in the array
// that make the specified condition true.
func (s Stream[T]) Filter(expression func(item T) bool) Stream[T] {
	return Stream[T]{
		Array: filter(expression, s.Array),
	}
}

// The Filter function leaves only those elements in the array
// that make the specified condition true.
func (s StructStream[T]) Filter(expression func(item T) bool) StructStream[T] {
	return StructStream[T]{
		Array: filter(expression, s.Array),
	}
}

func filter[T interface{}](expression func(item T) bool, arr []T) []T {
	res := []T{}
	for i := 0; i < len(arr); i++ {
		if expression(arr[i]) {
			res = append(res, arr[i])
		}
	}

	return res
}
