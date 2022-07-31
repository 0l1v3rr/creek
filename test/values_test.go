package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestValues(t *testing.T) {
	var tests = []struct {
		array    map[int]string
		expected []string
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			expected: []string{"Mark", "John", "Jack"},
		},
		{
			array:    map[int]string{4: "Mark", 5: "Johnn"},
			expected: []string{"Mark", "Johnn"},
		},
		{
			array:    map[int]string{1: "Oliver", 2: "Oliver"},
			expected: []string{"Oliver", "Oliver"},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Values(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).Values().OrderByDescending().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
