package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findShortestSubArray([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2})) //10
	t.Log(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	t.Log(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
