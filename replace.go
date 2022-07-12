package creek

import "reflect"

// The Replace function replaces every occurrence of 'from' to 'to'.
// The first parameter is 'from', and the second is 'to'.
func (s Stream[T]) Replace(from T, to T) Stream[T] {
	return Stream[T]{
		Array: replace(from, to, s.Array),
	}
}

// The Replace function replaces every occurrence of 'from' to 'to'.
// The first parameter is 'from', and the second is 'to'.
func (s StructStream[T]) Replace(from T, to T) StructStream[T] {
	return StructStream[T]{
		Array: replace(from, to, s.Array),
	}
}

// The ReplaceWhere function replaces every element that satisfies the condition.
func (s Stream[T]) ReplaceWhere(function func(item T) bool, to T) Stream[T] {
	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if function(s.Array[i]) {
			result = append(result, to)
			continue
		}

		result = append(result, s.Array[i])
	}

	return Stream[T]{
		Array: result,
	}
}

// The ReplaceWhere function replaces every element that satisfies the condition.
func (s StructStream[T]) ReplaceWhere(function func(item T) bool, to T) StructStream[T] {
	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if function(s.Array[i]) {
			result = append(result, to)
			continue
		}

		result = append(result, s.Array[i])
	}

	return StructStream[T]{
		Array: result,
	}
}

func replace[T interface{}](from T, to T, arr []T) []T {
	result := []T{}

	for i := 0; i < len(arr); i++ {
		if reflect.DeepEqual(arr[i], from) {
			result = append(result, to)
			continue
		}

		result = append(result, arr[i])
	}

	return result
}
