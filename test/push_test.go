package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestPush(t *testing.T) {
	var tests = []struct {
		array    []int
		toPush   []int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			toPush:   []int{8, 2, 5, 12, 9},
			expected: []int{2, 7, 3, 1, 9, 12, 5, 8, 2, 5, 12, 9},
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			toPush:   []int{},
			expected: []int{6, 82, 1, 3, 1343, 12, 9},
		},
		{
			array:    []int{},
			toPush:   []int{6, 82, 1, 3, 1343, 12, 9},
			expected: []int{6, 82, 1, 3, 1343, 12, 9},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Push(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Push(item.toPush).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestPushValues(t *testing.T) {
	testname := "PushValues(): #1"

	t.Run(testname, func(t *testing.T) {
		result := creek.FromArray([]int{1, 8, 2, 7}).PushValues(7, 2).Collect()
		expected := []int{1, 8, 2, 7, 7, 2}
		if reflect.DeepEqual(result, expected) {
			t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, expected, result)
			return
		}

		t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, expected, result)
	})
}
