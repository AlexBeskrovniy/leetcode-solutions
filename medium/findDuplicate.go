package medium

func findDuplicate(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = m[num] + 1
		if m[num] > 1 {
			return num
		}
	}

	return 0
}

//Try to optimize - actually gives the worse result than the first solution
// func findDuplicate(nums []int) int {
// 	m := make(map[int]bool)

// 	for _, num := range nums {
// 		if m[num] {
// 			return num
// 		}

// 		m[num] = true
// 	}

// 	return -1
// }
