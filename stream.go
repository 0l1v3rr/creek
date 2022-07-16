package creek

import (
	"bufio"
	"os"
	"reflect"
)

const ByKey bool = true
const ByValue bool = false

// The Streamable interface defines every type you can use the streams with.
type Streamable interface {
	string | byte | float32 | float64 | int | int16 | int32 | int64 | uint16 | uint32 | uint64
}

type StructStream[T interface{}] struct {
	Array []T
}

type Stream[T Streamable] struct {
	Array []T
}

type MapStream[T Streamable, V Streamable] struct {
	Array []KeyValuePair[T, V]
}

type KeyValuePair[T Streamable, V Streamable] struct {
	Key   T
	Value V
}

// The FromMap function creates a new stream from the given map.
func FromMap[T Streamable, V Streamable](m map[T]V) MapStream[T, V] {
	arr := []KeyValuePair[T, V]{}

	for key, element := range m {
		arr = append(arr, KeyValuePair[T, V]{
			Key:   key,
			Value: element,
		})
	}

	return MapStream[T, V]{
		Array: arr,
	}
}

// The EmptyMap function creates a new empty stream for maps.
func EmptyMap[T Streamable, V Streamable]() MapStream[T, V] {
	return MapStream[T, V]{
		Array: []KeyValuePair[T, V]{},
	}
}

// The FromArray function creates a new stream from the given array.
func FromArray[T Streamable](array []T) Stream[T] {
	return Stream[T]{
		Array: array,
	}
}

// The FromStruct function creates a new stream from the given struct array.
// If the given array is not made of struct, it throws an error.
func FromStructs[T interface{}](array []T) StructStream[T] {
	for i := 0; i < len(array); i++ {
		if reflect.ValueOf(array[i]).Kind() != reflect.Struct {
			panic("the passed array is not made of structs")
		}
	}

	return StructStream[T]{
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

// The FromFile function creates a stream from a file.
// The file is read line by line. Each line is an element of the stream.
func FromFile(file *os.File) Stream[string] {
	scanner := bufio.NewScanner(file)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()
	return Stream[string]{
		Array: result,
	}
}

// The Collect function returns the modified array from the streams.
func (s Stream[T]) Collect() []T {
	return s.Array
}

// The Collect function returns the modified array from the streams.
func (s StructStream[T]) Collect() []T {
	return s.Array
}

// The Collect function returns the modified map from the streams.
func (s MapStream[T, V]) Collect() map[T]V {
	res := map[T]V{}

	for i := 0; i < len(s.Array); i++ {
		res[s.Array[i].Key] = s.Array[i].Value
	}

	return res
}
