package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2))
	t.Log(findWinningPlayer([]int{4, 2, 6, 3, 9}, 3))
	t.Log(findWinningPlayer([]int{2, 5, 4}, 3))
	t.Log(findWinningPlayer([]int{2, 5, 4}, 4))
	t.Log(findWinningPlayer([]int{16, 4, 7, 17}, 562084119))
}
