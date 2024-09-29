package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(timeRequiredToBuy([]int{2, 3, 2}, 2))
	t.Log(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
	t.Log(timeRequiredToBuy([]int{1, 3, 2, 4}, 1))
}
