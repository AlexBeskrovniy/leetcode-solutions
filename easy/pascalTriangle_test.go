package easy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneratePascal(t *testing.T) {
	numRows := 5
	expect := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}

	result := Generate(numRows)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expect - %v, Result - %v", expect, result)
	} else {
		fmt.Println("Passed")
	}
}

func TestGeneratePascalWithOne(t *testing.T) {
	numRows := 1
	expect := [][]int{{1}}

	result := Generate(numRows)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expect - %v, Result - %v", expect, result)
	} else {
		fmt.Println("Passed")
	}
}
