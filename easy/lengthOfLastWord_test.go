package easy

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	res := LengthOfLastWord("Hello World")

	if res != 5 {
		t.Error("Result of the LengthOfLastWord should be a 5")
	}
}

func TestLengthOfLastWordEmpty(t *testing.T) {
	res := LengthOfLastWord("")

	if res != 0 {
		t.Error("Result of the LengthOfLastWord should be a 0")
	}
}
