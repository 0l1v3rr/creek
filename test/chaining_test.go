package test

import (
	"reflect"
	"testing"

	"github.com/0l1v3rr/creek"
)

func TestChaining(t *testing.T) {
	arr := []int{2, 7, 3, 1, 12, 6, 82, 101, 23, 24, 72, 13, 7}
	result := creek.FromArray(arr).Filter(func(item int) bool {
		return item%2 == 0
	}).Map(func(item int) int {
		return item * 3
	}).OrderByDescending().Limit(5).Collect()

	expected := []int{246, 216, 72, 36, 18}

	if reflect.DeepEqual(result, expected) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected, result)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected, result)
	}
}
