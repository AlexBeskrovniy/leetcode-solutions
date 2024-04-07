package easy

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	var short, long string
	var res string

	if len(a) > len(b) {
		short = b
		long = a
	} else {
		short = a
		long = b
	}

	diff := len(long) - len(short)
	for i := diff; i > 0; i-- {
		short = "0" + short
	}

	var complement int
	var prepend string
	for i := len(long) - 1; i >= 0; i-- {
		first, _ := strconv.Atoi(string(long[i]))
		second, _ := strconv.Atoi(string(short[i]))

		add := first + second + complement

		if add > 1 {
			if add == 2 {
				prepend = "0"
			} else {
				prepend = "1"
			}
			complement = 1
		} else {
			if add == 0 {
				prepend = "0"
			} else {
				prepend = "1"
			}
			complement = 0
		}
		res = prepend + res

		if i == 0 {
			if complement != 0 {
				res = "1" + res
			}
		}
	}
	return res
}
