package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestElementAt(t *testing.T) {
	var tests = []struct {
		array    []int
		index    int
		expected int
	}{
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    2,
			expected: 1,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    5,
			expected: 9,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    0,
			expected: 3,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).ElementAt(item.index)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestElementAtOrElse(t *testing.T) {
	var tests = []struct {
		array    []int
		index    int
		orElse   int
		expected int
	}{
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    2,
			orElse:   101,
			expected: 1,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    5,
			orElse:   101,
			expected: 9,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9},
			index:    101,
			orElse:   101,
			expected: 101,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).ElementAtOrElse(item.index, item.orElse)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestFirst(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{9, 3, 4, 1, 4, 2, 9, 7},
			expected: 9,
		},
		{
			array:    []int{2, 3, 4, 1, 4, 2, 9},
			expected: 2,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9, 43},
			expected: 3,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).First()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestLast(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{9, 3, 4, 1, 4, 2, 9, 7},
			expected: 7,
		},
		{
			array:    []int{2, 3, 4, 1, 4, 2, 9},
			expected: 9,
		},
		{
			array:    []int{3, 4, 1, 4, 2, 9, 43},
			expected: 43,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Last()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
