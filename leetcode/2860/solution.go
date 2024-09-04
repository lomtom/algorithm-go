package leetcode

import "sort"

// 执行耗时:76 ms,击败了53.33% 的Go用户
// 内存消耗:8.3 MB,击败了13.33% 的Go用户
func countWays(nums []int) int {
	sort.Ints(nums)
	// 边界 下标：len(nums) - 1
	var res int = 1
	// 边界 下标：0
	if nums[0] != 0 {
		res++
	}
	for index := 0; index < len(nums)-1; index++ {
		if index+1 > nums[index] && index+1 < nums[index+1] {
			res++
		}
	}
	return res
}
