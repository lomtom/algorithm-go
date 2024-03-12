package leetcode

import "testing"

func Test(t *testing.T) {
	root := construct([]int{-1, -1, -1})
	e := Constructor(root)
	t.Log(e.Find(1))
	t.Log(e.Find(2))
}
