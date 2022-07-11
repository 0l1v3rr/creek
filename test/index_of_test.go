package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestIndexOf(t *testing.T) {
	var tests = []struct {
		array    []int
		toFind   int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 4},
			toFind:   2,
			expected: 0,
		},
		{
			array:    []int{2, 7, 3, 1, 4},
			toFind:   4,
			expected: 4,
		},
		{
			array:    []int{2, 7, 3, 1, 4, 3},
			toFind:   3,
			expected: 2,
		},
		{
			array:    []int{2, 7, 3, 1, 4, 3},
			toFind:   100,
			expected: -1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("IndexOf(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).IndexOf(item.toFind)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	var tests = []struct {
		array    []int
		toFind   int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 4},
			toFind:   2,
			expected: 0,
		},
		{
			array:    []int{2, 7, 3, 1, 4, 4},
			toFind:   4,
			expected: 5,
		},
		{
			array:    []int{2, 7, 3, 1, 4, 3},
			toFind:   3,
			expected: 5,
		},
		{
			array:    []int{2, 7, 3, 1, 4, 3},
			toFind:   100,
			expected: -1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("LastIndexOf(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).LastIndexOf(item.toFind)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
