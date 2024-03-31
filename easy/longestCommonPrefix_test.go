package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}

	output := LongestCommonPrefix(strs)

	if output != "fl" {
		t.Error("Result of LongestCommonPrefix should be a \"fl\"")
	}
}

func TestLongestCommonPrefixNoPrefix(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}

	output := LongestCommonPrefix(strs)

	if output != "" {
		t.Error("Result of LongestCommonPrefix should be an empty string")
	}
}

func TestLongestCommonPrefixNoPrefixZeroSlice(t *testing.T) {
	strs := []string{}

	output := LongestCommonPrefix(strs)

	if output != "" {
		t.Error("Result of LongestCommonPrefix should be an empty string")
	}
}

func TestLongestCommonPrefixNoPrefixEdgeCase(t *testing.T) {
	strs := []string{"cir", "car"}

	output := LongestCommonPrefix(strs)

	if output != "c" {
		t.Errorf("Result of LongestCommonPrefix should be a \"c\" but not %v", output)
	}
}
