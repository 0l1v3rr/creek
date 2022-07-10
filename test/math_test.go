package test

import (
	"fmt"
	"reflect"
	"strconv"
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

	// -----------------

	result2 := creek.FromStructs(GetOtherStructArray()).Max("Id")
	expected2 := TestStruct{Id: 14, Name: "Paul"}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}

	// -----------------

	result2 = creek.FromStructs(GetTestStructArray()).Max("Name")
	expected2 = TestStruct{Id: 2, Name: "Will"}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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

	// -----------------

	result2 := creek.FromStructs(GetOtherStructArray()).MaxIndex("Id")
	expected2 := 1

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}

	// -----------------

	result2 = creek.FromStructs(GetTestStructArray()).MaxIndex("Name")
	expected2 = 1

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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

	// -----------------

	result2 := creek.FromStructs(GetOtherStructArray()).Min("Id")
	expected2 := TestStruct{Id: 12, Name: "Ian"}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}

	// -----------------

	result2 = creek.FromStructs(GetTestStructArray()).Min("Name")
	expected2 = TestStruct{Id: 1, Name: "John"}

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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

	// -----------------

	result2 := creek.FromStructs(GetOtherStructArray()).MinIndex("Id")
	expected2 := 0

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}

	// -----------------

	result2 = creek.FromStructs(GetTestStructArray()).MinIndex("Name")
	expected2 = 0

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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

	// -----------------

	result2, _ := strconv.Atoi(fmt.Sprintf("%v", creek.FromStructs(GetTestStructArray()).Sum("Id")))
	expected2 := 6

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
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

	// -----------------

	result2, _ := strconv.Atoi(fmt.Sprintf("%v", creek.FromStructs(GetTestStructArray()).Average("Id")))
	expected2 := 2

	if reflect.DeepEqual(result2, expected2) {
		t.Logf("Map PASSED - Expected: %v, got: %v", expected2, result2)
	} else {
		t.Errorf("Map FAILED - Expected: %v, got: %v", expected2, result2)
	}
}
