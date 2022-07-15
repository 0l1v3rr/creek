package creek

import (
	"fmt"
	"reflect"
	"strings"
)

// The Reverse function reverses the stream.
func (s Stream[T]) Reverse() Stream[T] {
	result := []T{}

	for i := len(s.Array) - 1; i >= 0; i-- {
		result = append(result, s.Array[i])
	}

	return Stream[T]{
		Array: result,
	}
}

// The Reverse function reverses the stream.
func (s StructStream[T]) Reverse() StructStream[T] {
	result := []T{}

	for i := len(s.Array) - 1; i >= 0; i-- {
		result = append(result, s.Array[i])
	}

	return StructStream[T]{
		Array: result,
	}
}

// Join concatenates the elements of the stream to create a single string.
// The passed parameter is placed between the elements.
func (s Stream[T]) Join(separator string) string {
	arrLen := len(s.Array)
	if arrLen == 0 {
		return ""
	}

	if arrLen == 1 {
		return fmt.Sprintf("%v", s.Array[0])
	}

	n := len(separator) * (arrLen - 1)
	for i := 0; i < len(s.Array); i++ {
		n += len(fmt.Sprintf("%v", s.Array[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprintf("%v", s.Array[0]))

	for _, se := range s.Array[1:] {
		b.WriteString(separator)
		b.WriteString(fmt.Sprintf("%v", se))
	}

	return b.String()
}

// The Contains function checks whether the stream contains the passed item.
func (s Stream[T]) Contains(item T) bool {
	for i := 0; i < len(s.Array); i++ {
		if s.Array[i] == item {
			return true
		}
	}

	return false
}

// The IsEmpty function checks whether the stream is empty.
func (s Stream[T]) IsEmpty() bool {
	return len(s.Array) == 0
}

// The IsNotEmpty function checks whether the stream is not empty.
func (s Stream[T]) IsNotEmpty() bool {
	return len(s.Array) != 0
}

// The Clear function clears every element from the stream.
func (s Stream[T]) Clear() Stream[T] {
	return Stream[T]{
		Array: []T{},
	}
}

// The Count function returns the count of elements in the stream.
func (s Stream[T]) Count() int {
	return len(s.Array)
}

// The Contains function checks whether the stream contains the passed item.
func (s StructStream[T]) Contains(item T) bool {
	for i := 0; i < len(s.Array); i++ {
		if reflect.DeepEqual(s.Array[i], item) {
			return true
		}
	}

	return false
}

// The IsEmpty function checks whether the stream is empty.
func (s StructStream[T]) IsEmpty() bool {
	return len(s.Array) == 0
}

// The IsNotEmpty function checks whether the stream is not empty.
func (s StructStream[T]) IsNotEmpty() bool {
	return len(s.Array) != 0
}

// The Clear function clears every element from the stream.
func (s StructStream[T]) Clear() StructStream[T] {
	return StructStream[T]{
		Array: []T{},
	}
}

// The Count function returns the count of elements in the stream.
func (s StructStream[T]) Count() int {
	return len(s.Array)
}

// The ContainsKey function checks whether the stream contains an item with the passed key.
func (s MapStream[T, V]) ContainsKey(key T) bool {
	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].Key == key {
			return true
		}
	}

	return false
}

// The IsEmpty function checks whether the stream is empty.
func (s MapStream[T, V]) IsEmpty() bool {
	return len(s.Array) == 0
}

// The IsNotEmpty function checks whether the stream is not empty.
func (s MapStream[T, V]) IsNotEmpty() bool {
	return len(s.Array) != 0
}

// The Clear function clears every element from the stream.
func (s MapStream[T, V]) Clear() Stream[T] {
	return Stream[T]{
		Array: []T{},
	}
}

// The Count function returns the count of elements in the stream.
func (s MapStream[T, V]) Count() int {
	return len(s.Array)
}
