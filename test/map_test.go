package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestMap(t *testing.T) {
	result := creek.FromArray([]int{2, 7, 3, 1}).Map(func(item int) int {
		return item * 2
	}).Collect()

	expected := []int{4, 14, 6, 2}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
