package easy

import "fmt"

// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right order.
// The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

func PlusOne(digits []int) []int {
	edge := 9
	lastIndex := len(digits) - 1
	if digits[lastIndex] != edge {
		digits[lastIndex]++
	} else {
		edgeCounter := 0
		for i := range digits {
			if digits[i] == edge {
				edgeCounter++
			}
		}

		if edgeCounter == len(digits) {
			newDigits := make([]int, len(digits)+1)
			newDigits[0] = 1
			fmt.Println(newDigits)
			return newDigits
		} else {
			if digits[lastIndex-1] != edge {
				digits[lastIndex] = 0
				digits[lastIndex-1]++
			} else {
				for i := len(digits) - 1; i >= 0; i-- {
					if digits[i] == edge {
						digits[i] = 0
					} else {
						digits[i]++
						return digits
					}
				}
			}
		}
	}

	return digits
}
