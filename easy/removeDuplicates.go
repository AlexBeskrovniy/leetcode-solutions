package easy

import "fmt"

func RemoveDuplicates(nums []int) int {
	uniqCounter := 0
	// iterationCounter := 0

	tailCounter := 0

	// if len(nums) > 0 {
	// 	uniqCounter++
	// }

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			uniqCounter++
		}
		if nums[i] != nums[i+1] {
			continue
		}

		for k := i; k < len(nums)-1; k++ {
			if nums[k] == nums[i] {
				nums = append(nums[:k], nums[k+1:]...)
				nums = append(nums, 0)
				// tailCounter++
				fmt.Println(tailCounter)
			}
		}

		fmt.Println(nums[i], nums[i+1])
		// if nums[i] != nums[i+1] {
		// 	uniqCounter++
		// }

		// iterationCounter++
		fmt.Println(nums)
	}

	// if iterationCounter == 0 {
	// 	uniqCounter = len(nums)
	// } else if iterationCounter == 1 {
	// 	uniqCounter = 1
	// } else {
	// 	uniqCounter++
	// }

	return uniqCounter
}
