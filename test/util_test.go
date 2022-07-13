package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: []int{5, 12, 9, 1, 3, 7, 2},
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: []int{9, 12, 1343, 3, 1, 82, 6},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: []int{75, 81, 52, 73, 19, 5},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Reverse(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Reverse().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestJoin(t *testing.T) {
	var tests = []struct {
		array    []int
		joinChar string
		expected string
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			joinChar: ", ",
			expected: "2, 7, 3, 1, 9, 12, 5",
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			joinChar: "asd",
			expected: "6asd82asd1asd3asd1343asd12asd9",
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			joinChar: "",
			expected: "51973528175",
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Join(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Join(item.joinChar)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		array    []int
		contains int
		expected bool
	}{
		{
			array:    []int{2, 7, 3, 1},
			contains: 2,
			expected: true,
		},
		{
			array:    []int{2, 7, 3, 1},
			contains: 23114,
			expected: false,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Contains(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Contains(item.contains)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		array    []int
		expected bool
	}{
		{
			array:    []int{},
			expected: true,
		},
		{
			array:    []int{2, 7, 3, 1},
			expected: false,
		},
		{
			array:    []int{2},
			expected: false,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("IsEmpty(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).IsEmpty()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{},
			expected: []int{},
		},
		{
			array:    []int{2, 7, 3, 1},
			expected: []int{},
		},
		{
			array:    []int{2},
			expected: []int{},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Clear(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Clear().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestCount(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{},
			expected: 0,
		},
		{
			array:    []int{2, 7, 3, 1},
			expected: 4,
		},
		{
			array:    []int{2},
			expected: 1,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Count(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Count()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
