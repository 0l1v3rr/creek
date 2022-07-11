package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestSome(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected bool
	}{
		{
			array:    []int{2, 7, 3, 1},
			expected: true,
			function: func(item int) bool {
				return item%2 == 0
			},
		},
		{
			array:    []int{3, 9, 5, 13},
			expected: false,
			function: func(item int) bool {
				return item%2 == 0
			},
		},
		{
			array:    []int{3, 9, 5, 13},
			expected: true,
			function: func(item int) bool {
				return item%3 == 0
			},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Some(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Some(item.function)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
