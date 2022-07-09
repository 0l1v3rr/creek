package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestMax(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).Max()
	expected := 12

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestMaxIndex(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).MaxIndex()
	expected := 5

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestMin(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).Min()
	expected := 1

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestMinIndex(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).MinIndex()
	expected := 3

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestSum(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).Sum()
	expected := 39

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}

func TestAverage(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1, 9, 12, 5}).Average()
	expected := 5.571428571428571

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
