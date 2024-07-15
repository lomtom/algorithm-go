package leetcode

import "sort"

// 执行耗时:165 ms,击败了90.48% 的Go用户
// 内存消耗:8.2 MB,击败了47.62% 的Go用户
func minPairSum(nums []int) int {
	sort.Ints(nums)
	max := 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[i]+nums[len(nums)-1-i] > max {
			max = nums[i] + nums[len(nums)-1-i]
		}
	}
	return max
}
