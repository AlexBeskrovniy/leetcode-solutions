package medium

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := 0
	h := len(nums)

	for l < h {
		m := (l + h) / 2

		if m == 0 && nums[m] < nums[m+1] {
			return nums[m]
		}

		if m == len(nums)-1 && nums[m] < nums[m-1] {
			return nums[m]
		}

		if nums[m] < nums[m-1] || nums[m] > nums[m+1] {
			return nums[m]
		}

		if nums[m-1] < nums[m] && nums[m] > nums[m+1] {

		}

	}

	return -1
}
