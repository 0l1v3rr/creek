package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestIndexOf(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).IndexOf(4)
	expected := 1

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestLastIndexOf(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).LastIndexOf(4)
	expected := 3

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
