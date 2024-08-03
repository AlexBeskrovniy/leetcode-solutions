package medium

func search(nums []int, n int) int {
	if n == 0 || n == len(nums)-1 {
		return n
	}

	if nums[n] > nums[n-1] && nums[n] > nums[n+1] {
		return n
	}

	var next int

	if nums[n-1] > nums[n] {
		next = n - 1
	}

	if nums[n+1] > nums[n] {
		next = n + 1
	}

	return search(nums, next)
}

func FindPeakElement(nums []int) int {
	if len(nums) > 1 {
		if len(nums) == 2 {
			if nums[1] > nums[0] {
				return 1
			} else {
				return 0
			}
		} else {
			l := 0
			h := len(nums) - 1
			n := (l + h) / 2

			return search(nums, n)
		}
	}

	return 0
}
