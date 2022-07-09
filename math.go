package creek

import (
	"fmt"
	"reflect"
	"strconv"
)

// The Max function returns the largest element from the stream.
func (s Stream[T]) Max() T {
	if reflect.TypeOf(s.Array[0]).Kind() == reflect.String {
		panic("the Max function doesn't work on string types")
	}

	max := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
	}

	return max
}

// The MaxIndex function returns the index of the largest element from the stream.
func (s Stream[T]) MaxIndex() int {
	if reflect.TypeOf(s.Array[0]).Kind() == reflect.String {
		panic("the MaxIndex function doesn't work on string types")
	}

	max := s.Array[0]
	maxIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
			maxIndex = i
		}
	}

	return maxIndex
}

// The Min function returns the smallest element from the stream.
func (s Stream[T]) Min() T {
	min := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] < min {
			min = s.Array[i]
		}
	}

	return min
}

// The MaxIndex function returns the index of the smallest element from the stream.
func (s Stream[T]) MinIndex() int {
	min := s.Array[0]
	minIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] < min {
			min = s.Array[i]
			minIndex = i
		}
	}

	return minIndex
}

// The Sum function adds up all values in a stream.
func (s Stream[T]) Sum() T {
	var sum T

	for i := 0; i < len(s.Array); i++ {
		sum += s.Array[i]
	}

	return sum
}

// The Average function calculates the average of the stream.
// This function doesn't work with strings.
func (s Stream[T]) Average() float64 {
	if reflect.TypeOf(s.Array[0]).Kind() == reflect.String {
		panic("the MinIndex function doesn't work on string types")
	}

	var sum T
	len := len(s.Array)

	for i := 0; i < len; i++ {
		sum += s.Array[i]
	}

	floatSum, _ := strconv.ParseFloat(fmt.Sprintf("%v", sum), 64)
	return floatSum / float64(len)
}
