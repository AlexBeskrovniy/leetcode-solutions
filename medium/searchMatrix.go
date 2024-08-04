package medium

func binarySearch(haystack []int, needle int) bool {
	l := 0
	h := len(haystack)

	for l < h {
		med := (l + h) / 2
		if haystack[med] == needle {
			return true
		}

		if haystack[med] < needle {
			l = med + 1
		} else {
			h = med
		}
	}

	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	for i := 0; i <= len(matrix)-1; i++ {
		if matrix[i][len(matrix[i])-1] >= target {
			if matrix[i][len(matrix[i])-1] == target {
				return true
			} else {
				return binarySearch(matrix[i], target)
			}
		}
	}

	return false
}
