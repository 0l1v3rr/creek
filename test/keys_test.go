package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestKeys(t *testing.T) {
	var tests = []struct {
		array    map[int]string
		expected []int
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			expected: []int{1, 2, 3},
		},
		{
			array:    map[int]string{4: "Mark", 5: "John", 6: "Jack"},
			expected: []int{4, 5, 6},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 4: "Jack"},
			expected: []int{1, 2, 4},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Keys(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).Keys().OrderBy().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
