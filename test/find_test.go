package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestFind(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected int
	}{
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: 2,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 2
			},
			expected: 7,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Find(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Find(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestFindIndex(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected int
	}{
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: 0,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 2
			},
			expected: 1,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 100
			},
			expected: -1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("FindIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).FindIndex(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestFindLast(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected int
	}{
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: 4,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 3
			},
			expected: 4,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("FindLast(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).FindLast(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestFindLastIndex(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected int
	}{
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: 4,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 2
			},
			expected: 4,
		},
		{
			array: []int{2, 7, 3, 1, 4},
			function: func(item int) bool {
				return item > 100
			},
			expected: -1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("FindLastIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).FindLastIndex(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
