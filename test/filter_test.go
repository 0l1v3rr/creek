package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestFilter(t *testing.T) {
	var tests = []struct {
		array    []int
		function func(item int) bool
		expected []int
	}{
		{
			array: []int{2, 7, 3, 1},
			function: func(item int) bool {
				return item > 2
			},
			expected: []int{7, 3},
		},
		{
			array: []int{2, 8, 9, 12, 7, 103},
			function: func(item int) bool {
				return item%2 == 0
			},
			expected: []int{2, 8, 12},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Filter(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Filter(item.function).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ------------

	var tests2 = []struct {
		array    []TestStruct
		function func(item TestStruct) bool
		expected []TestStruct
	}{
		{
			array: GetTestStructArray(),
			function: func(item TestStruct) bool {
				return item.Id > 2
			},
			expected: []TestStruct{{Id: 3, Name: "Mark"}},
		},
		{
			array: GetOtherStructArray(),
			function: func(item TestStruct) bool {
				return item.Id%2 == 0
			},
			expected: []TestStruct{{Id: 12, Name: "Ian"}, {Id: 14, Name: "Paul"}},
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("Filter(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).Filter(item.function).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
