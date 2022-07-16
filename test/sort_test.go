package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestOrderBy(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			expected: []int{1, 2, 3, 7},
		},
		{
			array:    []int{3, 9, 5, 13},
			expected: []int{3, 5, 9, 13},
		},
		{
			array:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("OrderBy(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).OrderBy().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// -----------

	var tests2 = []struct {
		array    []TestStruct
		expected []TestStruct
		field    string
	}{
		{
			array: []TestStruct{
				{Id: 1, Name: "John"},
				{Id: 3, Name: "Mark"},
				{Id: 2, Name: "Will"},
			},
			expected: []TestStruct{
				{Id: 1, Name: "John"},
				{Id: 2, Name: "Will"},
				{Id: 3, Name: "Mark"},
			},
			field: "Id",
		},
		{
			array: []TestStruct{
				{Id: 12, Name: "Ian"},
				{Id: 14, Name: "Paul"},
				{Id: 13, Name: "Josh"},
			},
			expected: []TestStruct{
				{Id: 12, Name: "Ian"},
				{Id: 13, Name: "Josh"},
				{Id: 14, Name: "Paul"},
			},
			field: "Name",
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("OrderBy(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).OrderBy(item.field).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// -----------

	var tests3 = []struct {
		array    map[int]string
		byKey    bool
		expected map[int]string
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			byKey:    true,
			expected: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 3: "Jack", 2: "John"},
			byKey:    true,
			expected: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			byKey:    false,
			expected: map[int]string{3: "Jack", 2: "John", 1: "Mark"},
		},
		{
			array:    map[int]string{3: "Jack", 2: "John", 1: "Mark"},
			byKey:    false,
			expected: map[int]string{3: "Jack", 2: "John", 1: "Mark"},
		},
	}

	for _, item := range tests3 {
		counter++
		testname := fmt.Sprintf("OrderBy(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).OrderBy(item.byKey).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestOrderByDescending(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{
			array:    []int{2, 7, 3, 1},
			expected: []int{7, 3, 2, 1},
		},
		{
			array:    []int{3, 9, 5, 13},
			expected: []int{13, 9, 5, 3},
		},
		{
			array:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("OrderByDescending(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).OrderByDescending().Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// -----------

	var tests2 = []struct {
		array    []TestStruct
		expected []TestStruct
		field    string
	}{
		{
			array: []TestStruct{
				{Id: 1, Name: "John"},
				{Id: 3, Name: "Mark"},
				{Id: 2, Name: "Will"},
			},
			expected: []TestStruct{
				{Id: 3, Name: "Mark"},
				{Id: 2, Name: "Will"},
				{Id: 1, Name: "John"},
			},
			field: "Id",
		},
		{
			array: []TestStruct{
				{Id: 12, Name: "Ian"},
				{Id: 14, Name: "Paul"},
				{Id: 13, Name: "Josh"},
			},
			expected: []TestStruct{
				{Id: 14, Name: "Paul"},
				{Id: 13, Name: "Josh"},
				{Id: 12, Name: "Ian"},
			},
			field: "Name",
		},
	}

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("OrderByDescending(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).OrderByDescending(item.field).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// -----------

	var tests3 = []struct {
		array    map[int]string
		byKey    bool
		expected map[int]string
	}{
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			byKey:    true,
			expected: map[int]string{3: "Jack", 2: "John", 1: "Mark"},
		},
		{
			array:    map[int]string{1: "Mark", 3: "Jack", 2: "John"},
			byKey:    true,
			expected: map[int]string{3: "Jack", 2: "John", 1: "Mark"},
		},
		{
			array:    map[int]string{1: "Mark", 2: "John", 3: "Jack"},
			byKey:    false,
			expected: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
		},
		{
			array:    map[int]string{3: "Jack", 2: "John", 1: "Mark"},
			byKey:    false,
			expected: map[int]string{1: "Mark", 2: "John", 3: "Jack"},
		},
	}

	for _, item := range tests3 {
		counter++
		testname := fmt.Sprintf("OrderByDescending(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromMap(item.array).OrderByDescending(item.byKey).Collect()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}
