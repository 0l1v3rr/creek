package creek

import "reflect"

// The Equals function compares two streams and returns true if they're equals.
func (s Stream[T]) Equals(b Stream[T]) bool {
	return s.ArrEquals(b.Array)
}

// The ArrEquals function compares the stream and the passed array and returns true if they're equals.
func (s Stream[T]) ArrEquals(arr []T) bool {
	sLen := len(s.Array)

	if sLen != len(arr) {
		return false
	}

	for i := 0; i < sLen; i++ {
		if s.Array[i] != arr[i] {
			return false
		}
	}

	return true
}

// The Equals function compares two streams and returns true if they're equals.
func (s StructStream[T]) Equals(b StructStream[T]) bool {
	return s.ArrEquals(b.Array)
}

// The ArrEquals function compares the stream and the passed array and returns true if they're equals.
func (s StructStream[T]) ArrEquals(arr []T) bool {
	sLen := len(s.Array)

	if sLen != len(arr) {
		return false
	}

	for i := 0; i < sLen; i++ {
		if !reflect.DeepEqual(s.Array[i], arr[i]) {
			return false
		}
	}

	return true
}
