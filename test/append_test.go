package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestAppend(t *testing.T) {
	var tests = []struct {
		array    []int
		toAppend int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			toAppend: 7,
			expected: []int{2, 7, 3, 1, 7},
		},
		{
			array:    []int{2, 8, 4, 12},
			toAppend: 9,
			expected: []int{2, 8, 4, 12, 9},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Append(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Append(item.toAppend).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestAppendAt(t *testing.T) {
	var tests = []struct {
		array       []int
		toAppend    int
		appendIndex int
		expected    []int
	}{
		{
			array:       []int{2, 7, 3, 1},
			toAppend:    7,
			appendIndex: 1,
			expected:    []int{2, 7, 7, 3, 1},
		},
		{
			array:       []int{2, 8, 4, 12},
			toAppend:    9,
			appendIndex: 0,
			expected:    []int{9, 2, 8, 4, 12},
		},
		{
			array:       []int{2, 8, 4, 12},
			toAppend:    9,
			appendIndex: 3,
			expected:    []int{2, 8, 4, 9, 12},
		},
		{
			array:       []int{2, 8, 4, 12},
			toAppend:    9,
			appendIndex: 4,
			expected:    []int{2, 8, 4, 12, 9},
		},
		{
			array:       []int{2, 8, 4, 12},
			toAppend:    9,
			appendIndex: 5,
			expected:    []int{2, 8, 4, 12},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("AppendAt(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).AppendAt(item.appendIndex, item.toAppend).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestAppendIf(t *testing.T) {
	var tests = []struct {
		array        []int
		toAppend     int
		shouldAppend bool
		expected     []int
	}{
		{
			array:        []int{2, 7, 3, 1},
			toAppend:     7,
			shouldAppend: true,
			expected:     []int{2, 7, 3, 1, 7},
		},
		{
			array:        []int{2, 8, 4, 12},
			toAppend:     9,
			shouldAppend: false,
			expected:     []int{2, 8, 4, 12},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("AppendIf(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).AppendIf(item.toAppend, item.shouldAppend).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
