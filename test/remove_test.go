package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestRemoveWhere(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected []int
	}{
		{
			array: []int{2, 7, 3, 1, 9, 12, 5},
			function: func(item int) bool {
				return item > 2
			},
			expected: []int{2, 1},
		},
		{
			array: []int{6, 82, 1, 3, 1343, 12, 9},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: []int{1, 3, 1343, 9},
		},
		{
			array: []int{5, 19, 73, 52, 81, 75},
			function: func(item int) bool {
				return item < 100
			},
			expected: []int{},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveWhere(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).RemoveWhere(item.function).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		array    []int
		toRemove int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			toRemove: 7,
			expected: []int{2, 3, 1, 9, 12, 5},
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			toRemove: 1343,
			expected: []int{6, 82, 1, 3, 12, 9},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			toRemove: 5,
			expected: []int{19, 73, 52, 81, 75},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			toRemove: 100,
			expected: []int{5, 19, 73, 52, 81, 75, 5},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Remove(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Remove(item.toRemove).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemoveAt(t *testing.T) {
	var tests = []struct {
		array    []int
		index    int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			index:    0,
			expected: []int{7, 3, 1, 9, 12, 5},
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			index:    3,
			expected: []int{6, 82, 1, 1343, 12, 9},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			index:    5,
			expected: []int{5, 19, 73, 52, 81, 5},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			index:    6,
			expected: []int{5, 19, 73, 52, 81, 75},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			index:    7,
			expected: []int{5, 19, 73, 52, 81, 75, 5},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveAt(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).RemoveAt(item.index).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: []int{2, 7, 3, 1, 9, 12, 5},
		},
		{
			array:    []int{6, 6, 6, 6, 6, 6},
			expected: []int{6},
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75, 5},
			expected: []int{5, 19, 73, 52, 81, 75},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveDuplicates(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).RemoveDuplicates().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemoveIf(t *testing.T) {
	var tests = []struct {
		array        []int
		toRemove     int
		shouldRemove bool
		expected     []int
	}{
		{
			array:        []int{2, 7, 3, 1, 9, 12, 5},
			toRemove:     7,
			shouldRemove: true,
			expected:     []int{2, 3, 1, 9, 12, 5},
		},
		{
			array:        []int{6, 82, 1, 3, 1343, 12, 9},
			toRemove:     1343,
			shouldRemove: true,
			expected:     []int{6, 82, 1, 3, 12, 9},
		},
		{
			array:        []int{5, 19, 73, 52, 81, 75, 5},
			toRemove:     5,
			shouldRemove: false,
			expected:     []int{5, 19, 73, 52, 81, 75, 5},
		},
		{
			array:        []int{5, 19, 73, 52, 81, 75, 5},
			toRemove:     100,
			shouldRemove: true,
			expected:     []int{5, 19, 73, 52, 81, 75, 5},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveIf(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).RemoveIf(item.toRemove, item.shouldRemove).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemoveKey(t *testing.T) {
	var tests = []struct {
		array    map[int]string
		toRemove int
		expected map[int]string
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: 1,
			expected: map[int]string{2: "John", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: 2,
			expected: map[int]string{1: "Mark", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: 3,
			expected: map[int]string{1: "Mark", 2: "John"},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveKey(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).RemoveKey(item.toRemove).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestRemoveValue(t *testing.T) {
	var tests = []struct {
		array    map[int]string
		toRemove string
		expected map[int]string
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: "Mark",
			expected: map[int]string{2: "John", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: "John",
			expected: map[int]string{1: "Mark", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			toRemove: "Jack",
			expected: map[int]string{1: "Mark", 2: "John"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack", 4: "Jack", 5: "Jack", 6: "Jack"},
			toRemove: "Jack",
			expected: map[int]string{1: "Mark", 2: "John"},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("RemoveValue(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).RemoveValue(item.toRemove).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
