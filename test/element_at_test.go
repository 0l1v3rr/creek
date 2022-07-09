package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestElementAt(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).ElementAt(2)
	expected := 1

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestElementAtOrElse(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).ElementAtOrElse(5, 100)
	expected := 9

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ----------------

	result = creek.FromArray([]int{3, 4, 1, 4, 2, 9}).ElementAtOrElse(6, 100)
	expected = 100

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
