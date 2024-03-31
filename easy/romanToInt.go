package easy

func RomanToInt(s string) int {
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
