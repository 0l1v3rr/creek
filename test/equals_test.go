package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestEquals(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Equals(creek.Stream[int]{
		Array: []int{2, 7, 3, 1},
	})
	expected := true

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ------------------

	result = creek.FromArray([]int{2, 7, 3, 1}).Equals(creek.Stream[int]{
		Array: []int{2, 7, 3, 2},
	})
	expected = false

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestArrEquals(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).ArrEquals([]int{2, 7, 3, 1})
	expected := true

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ------------------

	result = creek.FromArray([]int{2, 7, 3, 1}).ArrEquals([]int{2, 7, 3, 2})
	expected = false

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
