package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(shortestDistanceAfterQueries(7, [][]int{{0, 4}, {2, 6}}))
	t.Log(shortestDistanceAfterQueries(6, [][]int{{1, 3}, {3, 5}}))
	t.Log(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
	t.Log(shortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))
}
