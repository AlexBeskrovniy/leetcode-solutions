package easy

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3}

	output := PlusOne(digits)
	expect := []int{1, 2, 4}

	for i := range output {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
		fmt.Println(output)
	}
}

func TestAnotherPlusOne(t *testing.T) {
	digits := []int{4, 3, 2, 1}

	output := PlusOne(digits)
	expect := []int{4, 3, 2, 2}

	for i := range output {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}

func TestPlusOneWithNine(t *testing.T) {
	digits := []int{9, 9}

	output := PlusOne(digits)
	expect := []int{1, 0, 0}

	for i := range expect {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}

func TestPlusOneNotAllNines(t *testing.T) {
	digits := []int{4, 3, 9, 9}

	output := PlusOne(digits)
	expect := []int{4, 4, 0, 0}

	for i := range output {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}

func TestPlusOneEdgesNines(t *testing.T) {
	digits := []int{9, 8, 9}

	output := PlusOne(digits)
	expect := []int{9, 9, 0}

	for i := range expect {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}

func TestPlusOneEdgeCase(t *testing.T) {
	digits := []int{2, 4, 9, 3, 9}

	output := PlusOne(digits)
	expect := []int{2, 4, 9, 4, 0}

	for i := range expect {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}

func TestPlusAnotherOneEdgeCase(t *testing.T) {
	digits := []int{8, 9, 9, 9}

	output := PlusOne(digits)
	expect := []int{9, 0, 0, 0}

	for i := range expect {
		if output[i] != expect[i] {
			t.Errorf("Not incremented digit - %d", output[i])
		}
	}
	fmt.Println(output)
}
