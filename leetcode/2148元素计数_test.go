package leetcode

import (
	"sort"
	"testing"
)

func countElements(nums []int) (ans int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, num := range nums {
		if nums[0] > num && num > nums[len(nums)-1] {
			ans++
		}
	}
	return
}

func TestCountElements(t *testing.T) {
	countElements([]int{11, 7, 2, 15})
}
