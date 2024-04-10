package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maximumBinaryString("01111001100000110010")) // 11111111110111111111
	t.Log(maximumBinaryString("1100"))
	t.Log(maximumBinaryString("0011"))
	t.Log(maximumBinaryString("000110"))
	t.Log(maximumBinaryString("111110"))
	t.Log(maximumBinaryString("111111"))
	t.Log(maximumBinaryString("000000"))
	t.Log(maximumBinaryString("001"))
	t.Log(maximumBinaryString("01"))
}
