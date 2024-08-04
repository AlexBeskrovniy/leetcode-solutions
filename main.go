package main

import (
	"fmt"
	"leetcode/solutions/medium"
)

func main() {
	nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(medium.SearchMatrix(nums, target))
}
