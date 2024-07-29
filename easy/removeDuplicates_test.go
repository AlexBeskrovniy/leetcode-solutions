package easy

import "testing"

func TestRemoveDuplicatesBasic(t *testing.T) {
	nums := []int{1, 1, 2}
	expectNums := []int{1, 2}
	expectLen := 2

	uniqLen := RemoveDuplicates(nums)

	if expectLen == uniqLen {
		for i := 0; i < uniqLen; i++ {
			if nums[i] != expectNums[i] {
				t.Errorf("Invalid ansver")
			}
		}
	} else {
		t.Error("Invalid length of the uniq values")
	}
}

func TestRemoveDuplicatesNoNeedChange(t *testing.T) {
	nums := []int{1, 2}
	expectNums := []int{1, 2}
	expectLen := 2

	uniqLen := RemoveDuplicates(nums)

	if expectLen == uniqLen {
		for i := 0; i < uniqLen; i++ {
			if nums[i] != expectNums[i] {
				t.Errorf("Invalid ansver")
			}
		}
	} else {
		t.Error("Invalid length of the uniq values")
	}
}

func TestRemoveDuplicatesTailDuplication(t *testing.T) {
	nums := []int{1, 2, 2}
	expectNums := []int{1, 2}
	expectLen := 2

	uniqLen := RemoveDuplicates(nums)

	if expectLen == uniqLen {
		for i := 0; i < uniqLen; i++ {
			if nums[i] != expectNums[i] {
				t.Errorf("Invalid ansver")
			}
		}
	} else {
		t.Error("Invalid length of the uniq values")
	}
}

func TestRemoveDuplicatesTwoSameNums(t *testing.T) {
	nums := []int{1, 1}
	expectNums := []int{1}
	expectLen := 1

	uniqLen := RemoveDuplicates(nums)

	if expectLen == uniqLen {
		for i := 0; i < uniqLen; i++ {
			if nums[i] != expectNums[i] {
				t.Errorf("Invalid ansver")
			}
		}
	} else {
		t.Error("Invalid length of the uniq values")
	}
}
