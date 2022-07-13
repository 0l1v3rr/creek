package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		array    []int
		toSearch int
		expected int
	}{
		{
			array:    []int{1, 2, 4, 7},
			toSearch: 4,
			expected: 2,
		},
		{
			array:    []int{2, 7, 13, 93},
			toSearch: 13,
			expected: 2,
		},
		{
			array:    []int{0, 2, 7, 12},
			toSearch: 0,
			expected: 0,
		},
		{
			array:    []int{0, 2, 7, 12},
			toSearch: 12,
			expected: 3,
		},
		{
			array:    []int{0, 2, 7, 12},
			toSearch: 101,
			expected: -1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("BinarySearch(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).BinarySearch(item.toSearch)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
