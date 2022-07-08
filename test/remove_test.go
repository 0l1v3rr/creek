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
