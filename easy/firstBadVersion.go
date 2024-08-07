package easy

// You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version, all the versions after a bad version are also bad.
// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
// You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// Fake function just to check one test case
func isBadVersion(version int) bool {
	return version == 2
}

//My first try
func FirstBadVersion(n int) int {
	if n == 1 || !isBadVersion(n-1) {
		return n
	}

	l := 1
	h := n

	for l < h {
		m := (l + h) / 2

		if isBadVersion(m) {
			if m == 1 || !isBadVersion(m-1) {
				return m
			}

			h = m
		} else {
			l = m + 1
		}
	}

	return 1
}

//The same solution but more simple
func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
