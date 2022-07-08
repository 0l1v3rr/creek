package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestRemoveWhere(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).RemoveWhere(func(item int) bool {
		return item > 2
	}).Collect()

	expected := []int{2, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestRemove(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Remove(7).Collect()
	expected := []int{2, 3, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -------------------

	result = creek.FromArray([]int{2, 7, 3, 1, 3}).Remove(3).Collect()
	expected = []int{2, 7, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestRemoveAt(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).RemoveAt(2).Collect()
	expected := []int{2, 7, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 3, 9, 3}).RemoveDuplicates().Collect()
	expected := []int{2, 7, 3, 1, 9}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
