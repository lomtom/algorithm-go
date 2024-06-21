package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(temperatureTrend([]int{21, 18, 18, 18, 31}, []int{34, 32, 16, 16, 17}))
	t.Log(temperatureTrend([]int{5, 10, 16, -6, 15, 11, 3}, []int{16, 22, 23, 23, 25, 3, -16}))
}
