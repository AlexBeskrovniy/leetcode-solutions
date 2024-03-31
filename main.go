package main

import (
	"fmt"
	"leetcode/solutions/easy"
	"slices"
	"strconv"
	"strings"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
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

func isPalindrome(x int) bool {
	// Non optimized
	if x >= 0 {
		str := strconv.Itoa(x)
		slc := strings.Split(str, "")
		slices.Reverse(slc)
		if str == strings.Join(slc, "") {
			return true
		}

		return false
	}

	return false
}

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if i == 0 && target < num {
			return 0
		}
		if num == target {
			return i
		} else if i == len(nums)-1 {
			return i + 1
		} else if nums[i] < target && nums[i+1] > target {
			return i + 1
		}
	}

	return 0
}

func romanToInt(s string) int {
	numsMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var acc int

	for i := len(s) - 1; i >= 0; i-- {
		if i != len(s)-1 {
			if numsMap[string(s[i])] < numsMap[string(s[i+1])] {
				acc = acc - numsMap[string(s[i])]
			} else {
				acc = acc + numsMap[string(s[i])]
			}
		} else {
			acc = acc + numsMap[string(s[i])]
		}
	}
	// for i, char := range s {
	// 	if i != 0 {
	// 		if numsMap[string(char)] > numsMap[string(s[i-1])] {
	// 			acc = acc + numsMap[string(char)] - numsMap[string(s[i-1])]*2
	// 		} else {
	// 			acc = acc + numsMap[string(char)]
	// 		}

	// 	} else {
	// 		acc = acc + numsMap[string(char)]
	// 	}

	// }

	return acc
}

func main() {
	// str := "Hello world"
	// fmt.Println(easy.LengthOfLastWord(str))

	arr := []int{1, 1, 4, 4, 2}
	fmt.Println(easy.SingleNumber(arr))

	// arr := []int{2, 7, 11, 15}
	// target := 18

	// fmt.Println(twoSum(arr, target))

	// x := 221
	// fmt.Println(isPalindrome(x))

	// nums := []int{1, 3, 5, 6}
	// target := 0

	// fmt.Println(searchInsert(nums, target))

	// s := "MCMXCIV"
	// fmt.Println(romanToInt(s))
}
