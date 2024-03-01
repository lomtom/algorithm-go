package leetcode

// 2 ：2 2
// 3 ：2 2 2
// 3 ：2 3 4
// 执行耗时:108 ms,击败了47.62% 的Go用户
// 内存消耗:8.3 MB,击败了33.33% 的Go用户
func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := range nums {
		// 是为否：子数组 恰 由 2 个相等元素组成
		if i > 0 && dp[i-1] && nums[i] == nums[i-1] {
			dp[i+1] = true
		}
		// 是否为：子数组 恰 由 3 个相等元素组成
		if i > 1 && dp[i-2] && nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			dp[i+1] = true
		}
		// 是否为：子数组 恰 由 3 个连续递增元素组成
		if i > 1 && dp[i-2] && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			dp[i+1] = true
		}
	}
	return dp[n]
}
