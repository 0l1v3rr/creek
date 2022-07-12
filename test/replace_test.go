package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestReplace(t *testing.T) {
	var tests = []struct {
		array    []int
		from     int
		to       int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			from:     2,
			to:       3,
			expected: []int{3, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1, 3},
			from:     3,
			to:       2,
			expected: []int{2, 7, 2, 1, 2},
		},
		{
			array:    []int{2, 7, 3, 1},
			from:     5,
			to:       8,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{4, 4, 4, 4},
			from:     4,
			to:       4,
			expected: []int{4, 4, 4, 4},
		},
		{
			array:    []int{4, 4, 4, 4},
			from:     4,
			to:       3,
			expected: []int{3, 3, 3, 3},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Replace(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Replace(item.from, item.to).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestReplaceWhere(t *testing.T) {
	var tests = []struct {
		array    []int
		to       int
		function func(item int) bool
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			to:       3,
			expected: []int{3, 7, 3, 1},
			function: func(item int) bool {
				return item%2 == 0
			},
		},
		{
			array:    []int{2, 7, 3, 1, 3},
			to:       0,
			expected: []int{2, 0, 3, 1, 3},
			function: func(item int) bool {
				return item > 3
			},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("ReplaceWhere(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).ReplaceWhere(item.function, item.to).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
