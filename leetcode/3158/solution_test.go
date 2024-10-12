package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(duplicateNumbersXOR([]int{1, 2, 1, 3}))
	t.Log(duplicateNumbersXOR([]int{1, 2, 3}))
	t.Log(duplicateNumbersXOR([]int{1, 2, 2, 1}))
}
