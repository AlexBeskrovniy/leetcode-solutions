package easy

import (
	"fmt"
	"testing"
)

func TestAddBinary1(t *testing.T) {
	binaryStr1 := "11"
	binaryStr2 := "1"

	expect := "100"

	output := AddBinary(binaryStr1, binaryStr2)

	if output != expect {
		fmt.Printf("Output: %s\n Expect: %s", output, expect)
		t.Error("Invalid output")
	}
}

func TestAddBinary2(t *testing.T) {
	binaryStr1 := "1010"
	binaryStr2 := "1011"

	expect := "10101"

	output := AddBinary(binaryStr1, binaryStr2)

	if output != expect {
		fmt.Printf("Output: %s\n Expect: %s", output, expect)
		t.Error("Invalid output")
	}
}

func TestAddBinaryZero(t *testing.T) {
	binaryStr1 := "0"
	binaryStr2 := "0"

	expect := "0"

	output := AddBinary(binaryStr1, binaryStr2)

	if output != expect {
		fmt.Printf("Output: %s\n Expect: %s", output, expect)
		t.Error("Invalid output")
	}
}

func TestAddBinaryLong(t *testing.T) {
	binaryStr1 := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	binaryStr2 := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"

	expect := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"

	output := AddBinary(binaryStr1, binaryStr2)

	if output != expect {
		fmt.Printf("Output: %s\n Expect: %s", output, expect)
		t.Error("Invalid output")
	}
}

func TestAddBinaryOnlyOnes(t *testing.T) {
	binaryStr1 := "1111"
	binaryStr2 := "1111"

	expect := "11110"

	output := AddBinary(binaryStr1, binaryStr2)

	if output != expect {
		fmt.Printf("Output: %s\n Expect: %s", output, expect)
		t.Error("Invalid output")
	}
}
