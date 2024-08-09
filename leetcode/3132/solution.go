package leetcode

import "sort"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.7 MB,击败了63.64% 的Go用户
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result = nums2[0] - nums1[0]
	for i := 0; i <= 2; i++ {
		left, right := i+1, 1
		for left < len(nums1) && right < len(nums2) {
			if nums1[left]-nums2[right] == nums1[i]-nums2[0] {
				right++
			}
			left++
		}
		if right == len(nums2) && result > nums2[0]-nums1[i] {
			result = nums2[0] - nums1[i]
		}
	}
	return result
}
