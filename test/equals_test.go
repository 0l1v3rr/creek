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

	// ------------------

	result2 := creek.FromStructs(GetTestStructArray()).Equals(creek.FromStructs(GetTestStructArray()))
	expected2 := true

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}

	// ------------------

	result2 = creek.FromStructs(GetTestStructArray()).ArrEquals(GetOtherStructArray())
	expected2 = false

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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
