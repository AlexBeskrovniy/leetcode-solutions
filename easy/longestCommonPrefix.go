package easy

import (
	"sort"
)

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

func LongestCommonPrefix(strs []string) string {
	if len(strs) > 0 {
		if len(strs) == 1 {
			return strs[0]
		}
		sort.Slice(strs, func(i, j int) bool {
			return len(strs[i]) < len(strs[j])
		})

		var prefix string
		shortest := strs[0]
		for i, ch := range shortest {
			counter := 1
			for k := 1; k <= len(strs)-1; k++ {
				if rune(strs[k][i]) != ch {
					return prefix
				}
				counter++
			}
			if counter == len(strs) {
				prefix = prefix + string(ch)
			}
		}
		return prefix
	}
	return ""
}
