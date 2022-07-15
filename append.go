package creek

// The Append function adds an element to the stream.
func (s Stream[T]) Append(item T) Stream[T] {
	s.Array = append(s.Array, item)

	return Stream[T]{
		Array: s.Array,
	}
}

// The Append function adds an element to the stream.
func (s StructStream[T]) Append(item T) StructStream[T] {
	s.Array = append(s.Array, item)

	return StructStream[T]{
		Array: s.Array,
	}
}

// The Append function adds an element to the stream.
func (s MapStream[T, V]) Append(item KeyValuePair[T, V]) MapStream[T, V] {
	res := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].Key == item.Key {
			continue
		}

		res = append(res, s.Array[i])
	}

	res = append(res, item)

	return MapStream[T, V]{
		Array: res,
	}
}

// The AppendIf function adds an element to the stream if the second parameter is true.
func (s Stream[T]) AppendIf(item T, c bool) Stream[T] {
	if c {
		s.Array = append(s.Array, item)
	}

	return Stream[T]{
		Array: s.Array,
	}
}

// The AppendIf function adds an element to the stream if the second parameter is true.
func (s StructStream[T]) AppendIf(item T, c bool) StructStream[T] {
	if c {
		s.Array = append(s.Array, item)
	}

	return StructStream[T]{
		Array: s.Array,
	}
}

// The AppendIf function adds an element to the stream if the second parameter is true.
func (s MapStream[T, V]) AppendIf(item KeyValuePair[T, V], c bool) MapStream[T, V] {
	if c {
		s.Array = append(s.Array, item)
	}

	res := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].Key == item.Key {
			if c {
				continue
			}
		}

		res = append(res, s.Array[i])
	}

	if c {
		res = append(res, item)
	}

	return MapStream[T, V]{
		Array: res,
	}
}

// The AppendAt function inserts the specified element at the specified position in the stream.
func (s Stream[T]) AppendAt(index int, item T) Stream[T] {
	return Stream[T]{
		Array: appendAt(index, item, s.Array),
	}
}

// The AppendAt function inserts the specified element at the specified position in the stream.
func (s StructStream[T]) AppendAt(index int, item T) StructStream[T] {
	return StructStream[T]{
		Array: appendAt(index, item, s.Array),
	}
}

func appendAt[T interface{}](index int, item T, arr []T) []T {
	len := len(arr)
	if index == len {
		return append(arr, item)
	}

	if index > len {
		return arr
	}

	result := []T{}

	for i := 0; i < len; i++ {
		if i == index {
			result = append(result, item)
		}

		result = append(result, arr[i])
	}

	return result
}
