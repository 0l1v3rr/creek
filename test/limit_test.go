package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestLimit(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Limit(3).Collect()
	expected := []int{2, 7, 3}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -----------------------

	result = creek.FromArray([]int{2, 7, 3, 1}).Limit(7).Collect()
	expected = []int{2, 7, 3, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -----------------------

	result = creek.FromArray([]int{2, 7, 3, 1}).Limit(-3).Collect()
	expected = []int{}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
