package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(peakIndexInMountainArray([]int{0, 1, 0}))
	t.Log(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	t.Log(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	t.Log(peakIndexInMountainArray([]int{0, 10, 5, 2, 1}))
	t.Log(peakIndexInMountainArray([]int{0, 4, 10, 5, 2, 1}))
}
