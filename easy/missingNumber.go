package easy

func MissingNumber(nums []int) int {
	m := make(map[int]bool)

	for i := 0; i <= len(nums)-1; i++ {
		m[nums[i]] = true
	}

	for k := 0; k <= len(nums)+1; k++ {
		if !m[k] {
			return k
		}
	}

	return -1
}
