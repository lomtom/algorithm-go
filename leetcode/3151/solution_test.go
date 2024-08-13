package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(isArraySpecial([]int{1}))
	t.Log(isArraySpecial([]int{1, 2}))
	t.Log(isArraySpecial([]int{1, 2, 3, 1}))
}
