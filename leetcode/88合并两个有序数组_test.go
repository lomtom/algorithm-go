package leetcode

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums []int
	var index1, index2 = 0, 0
	for index1 < m && index2 < n {
		if nums1[index1] <= nums2[index2] {
			nums = append(nums, nums1[index1])
			index1++
		} else {
			nums = append(nums, nums2[index2])
			index2++
		}
	}
	nums = append(nums, nums1[index1:m]...)
	nums = append(nums, nums2[index2:]...)
	copy(nums1, nums)
}

func TestMerge(t *testing.T) {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}
