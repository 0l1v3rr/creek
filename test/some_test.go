package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestSome(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Some(func(item int) bool {
		return item%2 == 0
	})
	expected := true

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -----------------

	result = creek.FromArray([]int{3, 9, 5, 13}).All(func(item int) bool {
		return item%2 == 0
	})
	expected = false

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
