package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(antiConstruct(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})))
	t.Log(antiConstruct(buildTree([]int{1}, []int{1})))
}
