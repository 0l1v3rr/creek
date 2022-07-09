package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestFind(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 4}).Find(func(item int) bool {
		return item%2 == 0
	})
	expected := 2

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestFindIndex(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 4}).FindIndex(func(item int) bool {
		return item%2 == 0
	})
	expected := 0

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ----------------------

	result = creek.FromArray([]int{2, 7, 3, 1, 4}).FindIndex(func(item int) bool {
		return item > 100
	})
	expected = -1

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestFindLast(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 4}).FindLast(func(item int) bool {
		return item%2 == 0
	})
	expected := 4

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestFindLastIndex(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 4}).FindLastIndex(func(item int) bool {
		return item%2 == 0
	})
	expected := 4

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// ----------------------

	result = creek.FromArray([]int{2, 7, 3, 1, 4}).FindLastIndex(func(item int) bool {
		return item > 100
	})
	expected = -1

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
