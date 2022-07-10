package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestOrderBy(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).OrderBy().Collect()
	expected := []int{1, 2, 3, 7}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -----------

	result2 := creek.FromStructs(GetTestStructArray()).OrderBy("Name").Collect()
	expected2 := []TestStruct{
		{Id: 1, Name: "John"},
		{Id: 3, Name: "Mark"},
		{Id: 2, Name: "Will"},
	}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}
}

func TestOrderByDescending(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).OrderByDescending().Collect()
	expected := []int{7, 3, 2, 1}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}

	// -----------

	result2 := creek.FromStructs(GetTestStructArray()).OrderByDescending("Name").Collect()
	expected2 := []TestStruct{
		{Id: 2, Name: "Will"},
		{Id: 3, Name: "Mark"},
		{Id: 1, Name: "John"},
	}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}
}
