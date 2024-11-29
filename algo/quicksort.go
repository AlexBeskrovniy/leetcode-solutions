package algo

func QuickSort(nums []int) []int {
	var pivot int
	var lt []int
	var gt []int

	if len(nums) < 2 {
		return nums
	}

	pivot = nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}

		if num >= pivot {
			gt = append(gt, num)
		}

		if num < pivot {
			lt = append(lt, num)
		}
	}

	return append(append(QuickSort(lt), pivot), QuickSort(gt)...)
}
