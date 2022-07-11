package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestEquals(t *testing.T) {
	var tests = []struct {
		array    []int
		stream   creek.Stream[int]
		expected bool
	}{
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			stream:   creek.Stream[int]{Array: []int{3, 4, 1, 4, 2, 9}},
			expected: true,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			stream:   creek.Stream[int]{Array: []int{3, 4, 1, 4, 9}},
			expected: false,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Equals(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Equals(item.stream)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ------------------

	var tests2 = []struct {
		array    []TestStruct
		stream   creek.StructStream[TestStruct]
		expected bool
	}{
		{
			array:    GetTestStructArray(),
			stream:   creek.StructStream[TestStruct]{Array: GetTestStructArray()},
			expected: true,
		},
		{
			array:    GetTestStructArray(),
			stream:   creek.StructStream[TestStruct]{Array: GetOtherStructArray()},
			expected: false,
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("Equals(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).Equals(item.stream)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestArrEquals(t *testing.T) {
	var tests = []struct {
		array    []int
		stream   []int
		expected bool
	}{
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			stream:   []int{3, 4, 1, 4, 2, 9},
			expected: true,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			stream:   []int{3, 4, 1, 4, 9},
			expected: false,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("ArrEquals(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).ArrEquals(item.stream)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ------------------

	var tests2 = []struct {
		array    []TestStruct
		stream   []TestStruct
		expected bool
	}{
		{
			array:    GetTestStructArray(),
			stream:   GetTestStructArray(),
			expected: true,
		},
		{
			array:    GetTestStructArray(),
			stream:   GetOtherStructArray(),
			expected: false,
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("ArrEquals(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).ArrEquals(item.stream)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
