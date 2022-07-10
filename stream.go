package creek

import (
	"bufio"
	"os"
	"reflect"
)

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
