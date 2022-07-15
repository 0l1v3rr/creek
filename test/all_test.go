package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestAll(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected bool
	}{
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: false,
		},
		{
			array: []int{2, 8, 4, 12},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: true,
		},
		{
			array: []int{2, 8, 4, 8, 2, 4, 12, 9, 8, 22, 75},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: false,
		},
		{
			array: []int{2, 8, 4, 8, 2, 4, 12, 10, 8, 22, 74},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: true,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).All(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// MAP TESTS

	var tests2 = []struct {
		array    map[int]string
		function func(item creek.KeyValuePair[int, string]) bool
		expected bool
	}{
		{
			array: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			function: func(item creek.KeyValuePair[int, string]) bool {
				return item.Key > 0
			},
			expected: true,
		},
		{
			array: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			function: func(item creek.KeyValuePair[int, string]) bool {
				return len(item.Value) < 2
			},
			expected: false,
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).All(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
