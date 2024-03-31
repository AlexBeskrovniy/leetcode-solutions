package easy

import (
	"slices"
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	// Non optimized
	if x >= 0 {
		str := strconv.Itoa(x)
		slc := strings.Split(str, "")
		slices.Reverse(slc)
		return str == strings.Join(slc, "")
	}

	return false
}
