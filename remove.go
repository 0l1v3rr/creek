package creek

import "reflect"

// The RemoveWhere function removes all the entries that satisfy the provided condition.
func (s Stream[T]) RemoveWhere(expression func(item T) bool) Stream[T] {
	result := []T{}
	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			result = append(result, s.Array[i])
		}
	}

	return Stream[T]{
		Array: result,
	}
}

// The Remove function removes the passed element from a stream.
func (s Stream[T]) Remove(item T) Stream[T] {
	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if s.Array[i] != item {
			result = append(result, s.Array[i])
		}
	}

	return Stream[T]{
		Array: result,
	}
}

// The RemoveIf function removes the passed element from a stream if the second parameter is true.
func (s Stream[T]) RemoveIf(item T, c bool) Stream[T] {
	if !c {
		return s
	}

	return s.Remove(item)
}

// The RemoveAt function removes the item if its index matches the index passed in.
func (s Stream[T]) RemoveAt(index int) Stream[T] {
	if index < 0 {
		return s
	}

	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if i != index {
			result = append(result, s.Array[i])
		}
	}

	return Stream[T]{
		Array: result,
	}
}

// The RemoveDuplicates function removes every duplicate item from the stream.
func (s Stream[T]) RemoveDuplicates() Stream[T] {
	keys := make(map[T]bool)
	result := []T{}

	for _, item := range s.Array {
		if _, value := keys[item]; !value {
			keys[item] = true
			result = append(result, item)
		}
	}

	return Stream[T]{
		Array: result,
	}
}

// The Distinct function filters every distinct element from the stream.
func (s Stream[T]) Distinct() Stream[T] {
	return s.RemoveDuplicates()
}

// The RemoveWhere function removes all the entries that satisfy the provided condition.
func (s StructStream[T]) RemoveWhere(expression func(item T) bool) StructStream[T] {
	result := []T{}
	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			result = append(result, s.Array[i])
		}
	}

	return StructStream[T]{
		Array: result,
	}
}

// The Remove function removes the passed element from a stream.
func (s StructStream[T]) Remove(item T) StructStream[T] {
	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if !reflect.DeepEqual(s.Array[i], item) {
			result = append(result, s.Array[i])
		}
	}

	return StructStream[T]{
		Array: result,
	}
}

// The RemoveIf function removes the passed element from a stream if the second parameter is true.
func (s StructStream[T]) RemoveIf(item T, c bool) StructStream[T] {
	if !c {
		return s
	}

	return s.Remove(item)
}

// The RemoveAt function removes the item if its index matches the index passed in.
func (s StructStream[T]) RemoveAt(index int) StructStream[T] {
	if index < 0 {
		return s
	}

	result := []T{}

	for i := 0; i < len(s.Array); i++ {
		if i != index {
			result = append(result, s.Array[i])
		}
	}

	return StructStream[T]{
		Array: result,
	}
}

// The RemoveWhere function removes all the entries that satisfy the provided condition.
func (s MapStream[T, V]) RemoveWhere(expression func(KeyValuePair[T, V]) bool) MapStream[T, V] {
	result := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		if !expression(s.Array[i]) {
			result = append(result, s.Array[i])
		}
	}

	return MapStream[T, V]{
		Array: result,
	}
}

// The RemoveKey function removes every element from a stream
// where the key is equal to the passed item.
func (s MapStream[T, V]) RemoveKey(key T) MapStream[T, V] {
	result := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].Key == key {
			continue
		}

		result = append(result, s.Array[i])
	}

	return MapStream[T, V]{
		Array: result,
	}
}

// The RemoveValue function removes every element from a stream
// where the value is equal to the passed item.
func (s MapStream[T, V]) RemoveValue(value V) MapStream[T, V] {
	result := []KeyValuePair[T, V]{}

	for i := 0; i < len(s.Array); i++ {
		if s.Array[i].Value == value {
			continue
		}

		result = append(result, s.Array[i])
	}

	return MapStream[T, V]{
		Array: result,
	}
}
