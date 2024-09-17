package leetcode

import "sort"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了100.00% 的Go用户
func sumOfUnique(nums []int) int {
	sort.Ints(nums)
	ans := 0
	var left, right int
	for left < len(nums) {
		for right < len(nums) && nums[left] == nums[right] {
			right++
		}
		if right-left == 1 {
			ans += nums[left]
		}
		left = right
	}
	return ans
}
