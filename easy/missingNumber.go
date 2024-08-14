package easy

// func MissingNumber(nums []int) int {
// 	m := make(map[int]bool)

// 	for i := 0; i <= len(nums)-1; i++ {
// 		m[nums[i]] = true
// 	}

// 	for k := 0; k <= len(nums)+1; k++ {
// 		if !m[k] {
// 			return k
// 		}
// 	}

// 	return -1
// }

//More optimized way

func MissingNumber(nums []int) int {
	nums_sum := 0
	expected_sum := 0

	for i := 0; i <= len(nums); i++ {
		if i < len(nums) {
			nums_sum += nums[i]
		}

		expected_sum += i
	}

	return expected_sum - nums_sum
}
