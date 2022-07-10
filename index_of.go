package creek

import "reflect"

// The IndexOf function returns the position of the first occurrence of
// the passed value in a stream.
func (s Stream[T]) IndexOf(item T) int {
	for i := 0; i < len(s.Array); i++ {
		if s.Array[i] == item {
			return i
		}
	}

	return -1
}

// The LastIndexOf function returns the position of the last occurrence of
// the passed value in a stream.
func (s Stream[T]) LastIndexOf(item T) int {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if s.Array[i] == item {
			return i
		}
	}

	return -1
}

// The IndexOf function returns the position of the first occurrence of
// the passed value in a stream.
func (s StructStream[T]) IndexOf(item T) int {
	for i := 0; i < len(s.Array); i++ {
		if reflect.DeepEqual(s.Array[i], item) {
			return i
		}
	}

	return -1
}

// The LastIndexOf function returns the position of the last occurrence of
// the passed value in a stream.
func (s StructStream[T]) LastIndexOf(item T) int {
	for i := len(s.Array) - 1; i >= 0; i-- {
		if reflect.DeepEqual(s.Array[i], item) {
			return i
		}
	}

	return -1
}
