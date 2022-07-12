package creek

// The Slice function returns a copy of a portion of a stream
// into a new stream selected from start to end
// (end not included) where start and end represent the index of items in the stream.
func (s Stream[T]) Slice(start int, end int) Stream[T] {
	return Stream[T]{
		Array: slice(start, end, s.Array),
	}
}

// The Slice function returns a copy of a portion of a stream
// into a new stream selected from start to end
// (end not included) where start and end represent the index of items in the stream.
func (s StructStream[T]) Slice(start int, end int) StructStream[T] {
	return StructStream[T]{
		Array: slice(start, end, s.Array),
	}
}

func slice[T interface{}](start int, end int, arr []T) []T {
	if start > end {
		return arr
	}

	if start == end {
		return []T{arr[start-1]}
	}

	len := len(arr)

	if start >= len {
		return arr
	}

	if end > len {
		return arr[start:]
	}

	return arr[start:end]
}
