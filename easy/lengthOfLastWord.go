package easy

import "strings"

// Need to find lenght of the last word in a string
func LengthOfLastWord(s string) int {
	strSlice := strings.Fields(s)
	if len(strSlice) != 0 {
		return len(strSlice[len(strSlice)-1])
	}
	return 0
}
