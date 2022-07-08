package creek

// The Streamable interface defines every type you can use the streams with.
type streamable interface {
	string | byte | float32 | float64 | int | int16 | int32 | int64 | uint16 | uint32 | uint64
}

type Stream[T streamable] struct {
	Array []T
}

// The FromArray function creates a new stream from the given array.
func FromArray[T streamable](array []T) Stream[T] {
	return Stream[T]{
		Array: array,
	}
}

// The Collect function returns the modified array from the streams.
func (s Stream[T]) Collect() []T {
	return s.Array
}
