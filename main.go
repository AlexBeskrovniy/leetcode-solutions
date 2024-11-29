package main

import (
	"fmt"
	"leetcode/solutions/algo"
)

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1, 4, 4, 4, 8, 34, 64, 0, 2, 11, 5}
	fmt.Println(algo.QuickSort(nums))
}
