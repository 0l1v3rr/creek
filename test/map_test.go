package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestMap(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) int
		expected []int
	}{
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) int {
				return item * 2
			},
			expected: []int{4, 14, 6, 2},
		},
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) int {
				return item * item
			},
			expected: []int{4, 49, 9, 1},
		},
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) int {
				return item - 1
			},
			expected: []int{1, 6, 2, 0},
		},
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) int {
				return item * 0
			},
			expected: []int{0, 0, 0, 0},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Map(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Map(item.function).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
