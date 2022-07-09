package creek

// The Streamable interface defines every type you can use the streams with.
type Streamable interface {
	string | byte | float32 | float64 | int | int16 | int32 | int64 | uint16 | uint32 | uint64
}

type Stream[T Streamable] struct {
	Array []T
}

// The FromArray function creates a new stream from the given array.
func FromArray[T Streamable](array []T) Stream[T] {
	return Stream[T]{
		Array: array,
	}
}

// The Empty function returns an empty stream.
func Empty[T Streamable]() Stream[T] {
	return Stream[T]{
		Array: []T{},
	}
}

// The FromValues function returns a stream made of the specified parameters.
func FromValues[T Streamable](values ...T) Stream[T] {
	return Stream[T]{
		Array: values,
	}
}

// The Collect function returns the modified array from the streams.
func (s Stream[T]) Collect() []T {
	return s.Array
}
