package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestSlice(t *testing.T) {
	var tests = []struct {
		array    []int
		start    int
		end      int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			start:    0,
			end:      2,
			expected: []int{2, 7},
		},
		{
			array:    []int{2, 7, 3, 1},
			start:    3,
			end:      2,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			start:    5,
			end:      8,
			expected: []int{2, 7, 3, 1},
		},
		{
			array:    []int{2, 7, 3, 1},
			start:    4,
			end:      4,
			expected: []int{1},
		},
		{
			array:    []int{2, 7, 3, 1},
			start:    3,
			end:      3,
			expected: []int{3},
		},
		{
			array:    []int{2, 7, 3, 1},
			start:    3,
			end:      10,
			expected: []int{1},
		},
		{
			array:    []int{2, 7, 3, 1, 6, 3, 4, 12, 9, 13, 73, 17},
			start:    4,
			end:      10,
			expected: []int{6, 3, 4, 12, 9, 13},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Slice(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Slice(item.start, item.end).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
