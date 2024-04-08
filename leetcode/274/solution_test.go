package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(hIndex([]int{0}))
	t.Log(hIndex([]int{1}))
	t.Log(hIndex([]int{4}))
	t.Log(hIndex([]int{1000}))
	t.Log(hIndex([]int{1, 3, 1, 2}))
	t.Log(hIndex([]int{3, 0, 6, 1, 5}))
	t.Log(hIndex([]int{1, 3, 1}))
}
