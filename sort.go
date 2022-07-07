package creek

import "sort"

// The OrderBy function sorts the stream in ascending order.
func (s Stream[T]) OrderBy() Stream[T] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		return s.Array[i] < s.Array[j]
	})

	return Stream[T]{
		Array: s.Array,
	}
}

// The OrderByDescending function sorts the stream in descending order.
func (s Stream[T]) OrderByDescending() Stream[T] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		return s.Array[i] > s.Array[j]
	})

	return Stream[T]{
		Array: s.Array,
	}
}
