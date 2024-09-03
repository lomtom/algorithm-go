package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxStrength([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	t.Log(maxStrength([]int{-9}))
	t.Log(maxStrength([]int{0, -1}))
	t.Log(maxStrength([]int{3, -1, -5, 2, 5, -9}))
	t.Log(maxStrength([]int{-4, -5, -4}))
}
