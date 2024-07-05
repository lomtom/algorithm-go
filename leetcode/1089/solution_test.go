package leetcode

import "testing"

func Test(t *testing.T) {
	nums := []int{1, 0, 0, 3, 0, 4, 5, 0}
	duplicateZeros(nums)
	t.Log(nums)
	nums = []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(nums)
	t.Log(nums)
	nums = []int{1, 2, 3}
	duplicateZeros(nums)
	t.Log(nums)
	nums = []int{1, 0, 3}
	duplicateZeros(nums)
	t.Log(nums)
	nums = []int{0, 1, 3}
	duplicateZeros(nums)
	t.Log(nums)
	nums = []int{1, 1, 0}
	duplicateZeros(nums)
	t.Log(nums)
}
