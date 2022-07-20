package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestSkip(t *testing.T) {
	var tests = []struct {
		array    []int
		toSkip   int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   3,
			expected: []int{1},
		},
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   7,
			expected: []int{},
		},
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   -3,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   0,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   1,
			expected: []int{7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			toSkip:   4,
			expected: []int{},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Skip(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Skip(item.toSkip).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
