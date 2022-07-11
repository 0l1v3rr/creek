package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 12,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 1343,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 81,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Max(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Max()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ----------------------

	var tests2 = []struct {
		array    []TestStruct
		filed    string
		expected TestStruct
	}{
		{
			array:    GetTestStructArray(),
			expected: TestStruct{Id: 3, Name: "Mark"},
			filed:    "Id",
		},
		{
			array:    GetOtherStructArray(),
			expected: TestStruct{Id: 14, Name: "Paul"},
			filed:    "Id",
		},
	}

	// -----------------

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("Max(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).Max(item.filed)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestMaxIndex(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 5,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 4,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 4,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("MaxIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).MaxIndex()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ----------------------

	var tests2 = []struct {
		array    []TestStruct
		filed    string
		expected int
	}{
		{
			array:    GetTestStructArray(),
			expected: 2,
			filed:    "Id",
		},
		{
			array:    GetOtherStructArray(),
			expected: 1,
			filed:    "Id",
		},
	}

	// -----------------

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("MaxIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).MaxIndex(item.filed)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 1,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 1,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 5,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Min(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Min()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ----------------------

	var tests2 = []struct {
		array    []TestStruct
		filed    string
		expected TestStruct
	}{
		{
			array:    GetTestStructArray(),
			expected: TestStruct{Id: 1, Name: "John"},
			filed:    "Id",
		},
		{
			array:    GetOtherStructArray(),
			expected: TestStruct{Id: 12, Name: "Ian"},
			filed:    "Id",
		},
	}

	// -----------------

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("Min(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).Min(item.filed)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestMinIndex(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 3,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 2,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 0,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("MinIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).MinIndex()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ----------------------

	var tests2 = []struct {
		array    []TestStruct
		filed    string
		expected int
	}{
		{
			array:    GetTestStructArray(),
			expected: 0,
			filed:    "Id",
		},
		{
			array:    GetOtherStructArray(),
			expected: 0,
			filed:    "Id",
		},
	}

	// -----------------

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("MinIndex(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromStructs(item.array).MinIndex(item.filed)
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestSum(t *testing.T) {
	var tests = []struct {
		array    []int
		expected int
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 39,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 1456,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 305,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Sum(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Sum()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// ----------------------

	var tests2 = []struct {
		array    []TestStruct
		filed    string
		expected int
	}{
		{
			array:    GetTestStructArray(),
			expected: 6,
			filed:    "Id",
		},
		{
			array:    GetOtherStructArray(),
			expected: 39,
			filed:    "Id",
		},
	}

	// -----------------

	for _, item := range tests2 {
		counter++
		testname := fmt.Sprintf("Sum(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result, _ := strconv.Atoi(fmt.Sprintf("%v", creek.FromStructs(item.array).Sum(item.filed)))
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}
}

func TestAverage(t *testing.T) {
	var tests = []struct {
		array    []int
		expected float64
	}{
		{
			array:    []int{2, 7, 3, 1, 9, 12, 5},
			expected: 5.571428571428571,
		},
		{
			array:    []int{6, 82, 1, 3, 1343, 12, 9},
			expected: 208,
		},
		{
			array:    []int{5, 19, 73, 52, 81, 75},
			expected: 50.833333333333336,
		},
	}

	counter := 0

	for _, item := range tests {
		counter++
		testname := fmt.Sprintf("Average(): #%v", counter)

		t.Run(testname, func(t *testing.T) {
			result := creek.FromArray(item.array).Average()
			if reflect.DeepEqual(result, item.expected) {
				t.Logf("%v -> PASSED - Expected: %v, got: %v", testname, item.expected, result)
				return
			}

			t.Errorf("%v -> FAILED - Expected: %v, got: %v", testname, item.expected, result)
		})
	}

	// -----------------

	result2, _ := strconv.Atoi(fmt.Sprintf("%v", creek.FromStructs(GetTestStructArray()).Average("Id")))
	expected2 := 2

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Average PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Average FAILED - Expected: %v, got: %v", expected2, result2)
	}
}
