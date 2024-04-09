package leetcode

// 执行耗时:34 ms,击败了14.76% 的Go用户
// 内存消耗:8.1 MB,击败了10.55% 的Go用户
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	prefix, suffix := 1, 1
	for i := 0; i < n; i++ {
		ans[i] *= prefix
		ans[n-i-1] *= suffix
		prefix *= nums[i]
		suffix *= nums[n-i-1]
	}
	return ans
}
