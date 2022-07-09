package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestAppend(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Append(2).Collect()
	expected := []int{2, 7, 3, 1, 2}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestAppendAt(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).AppendAt(1, 14).Collect()
	expected := []int{2, 14, 7, 3, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestAppendIf(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).AppendIf(37, true).Collect()
	expected := []int{2, 7, 3, 1, 37}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ------------------

	result = creek.FromArray([]int{2, 7, 3, 1}).AppendIf(37, false).Collect()
	expected = []int{2, 7, 3, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
