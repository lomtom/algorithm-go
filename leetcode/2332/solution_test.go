package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(latestTimeCatchTheBus([]int{2}, []int{2}, 2))
	t.Log(latestTimeCatchTheBus([]int{10}, []int{2}, 2))
	t.Log(latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	t.Log(latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
}
