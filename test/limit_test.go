package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestLimit(t *testing.T) {
	var tests = []struct {
		array    []int
		toLimit  int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			toLimit:  3,
			expected: []int{2, 7, 3},
		},
		{
			array:    []int{2, 7, 3, 1},
			toLimit:  7,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			toLimit:  -3,
			expected: []int{},
		},
		{
			array:    []int{2, 7, 3, 1},
			toLimit:  0,
			expected: []int{},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("All(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Limit(item.toLimit).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
