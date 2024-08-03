package easy

func SearchRange(nums []int, target int) []int {
	var result = []int{-1, -1}

	if len(nums) != 0 {
		if len(nums) == 1 {
			if nums[0] == target {
				result[0] = 0
				result[1] = 0
			}
			return result
		}
		var l = 0
		var h = len(nums)
		var m int

		for l < h {
			m = (l + h) / 2
			if nums[m] == target {
				for i, k := m, m; i >= 0 || k <= len(nums); i, k = i-1, k+1 {
					if result[0] != -1 && result[1] != -1 {
						return result
					}
					if i > 0 {
						if nums[i-1] < target {
							if result[0] == -1 {
								result[0] = i
							}
						}
					} else {
						if nums[0] == target {
							result[0] = 0
						}
					}

					if k < len(nums)-1 {
						if nums[k+1] > target {
							if result[1] == -1 {
								result[1] = k
							}
						}
					} else {
						if nums[len(nums)-1] == target {
							result[1] = len(nums) - 1
						}
					}
				}
				return result
			} else {
				if nums[m] < target {
					l = m + 1
				} else {
					h = m
				}
			}
		}
	}

	return result
}
