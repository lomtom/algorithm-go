package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})) // [-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]
	t.Log(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
	t.Log(threeSum([]int{1, 2, -2, -1}))
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(threeSum([]int{0, 0, 0}))
	t.Log(threeSum([]int{0, 0, 0, 0}))
}
