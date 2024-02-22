package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(antiConstruct(constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})))
	t.Log(antiConstruct(constructFromPrePost([]int{1}, []int{1})))
}
