package easy

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
func SearchInsert(nums []int, target int) int {
	for i, num := range nums {
		if i == 0 && target < num {
			return 0
		}
		if num == target {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		} else if nums[i] < target && nums[i+1] > target {
			return i + 1
		}
	}

	return 0
}
