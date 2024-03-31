package easy

import "slices"

// a non-empty array of integers nums, every element appears twice except for one. Find that single one.
func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for i, value := range nums {
		if i == len(nums)-1 {
			return nums[i]
		}

		// Immutable copy of slice
		origin := make([]int, len(nums))
		copy(origin, nums)

		splicedArr := slices.Delete(origin, i, i+1)
		if !slices.Contains(splicedArr, value) {
			return nums[i]
		}
	}
	return 0
}
