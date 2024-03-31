package easy

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	// More optimized solution
	numsMap := make(map[int]int)

	for i, num := range nums {
		index, ok := numsMap[target-num]

		if ok {
			return []int{index, i}
		}
		numsMap[num] = i
	}
	return []int{0, 0}

	// Non optimized solution
	// for i, value := range nums {
	// 	for k, val := range nums {
	// 		if k == i {
	// 			continue
	// 		}

	// 		if target == val+value {
	// 			return []int{i, k}
	// 		}
	// 	}
	// }
	// return []int{0, 0}
}
