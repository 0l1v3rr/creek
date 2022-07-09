package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestReverse(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Reverse().Collect()
	expected := []int{1, 3, 7, 2}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestJoin(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Join(", ")
	expected := "2, 7, 3, 1"

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// --------------

	result = creek.FromArray([]int{2, 7, 3, 1}).Join(" asd, ")
	expected = "2 asd, 7 asd, 3 asd, 1"

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestContains(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Contains(2)
	expected := true

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// --------------

	result = creek.FromArray([]int{2, 7, 3, 1}).Contains(1213231)
	expected = false

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestIsEmpty(t *testing.T) {
	result := creek.FromArray([]int{}).IsEmpty()
	expected := true

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// --------------

	result = creek.FromArray([]int{2, 7, 3, 1}).IsEmpty()
	expected = false

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestClear(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).Clear().Collect()
	expected := []int{}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestCount(t *testing.T) {
	result := creek.FromArray([]int{3, 4, 1, 4, 2, 9}).Count()
	expected := 6

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
