package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(reachableNodes(2, [][]int{
		{0, 1}}, []int{1}))
	t.Log(reachableNodes(7, [][]int{
		{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}, []int{4, 2, 1}))
	t.Log(reachableNodes(7, [][]int{
		{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6},
	}, []int{4, 5}))
}
