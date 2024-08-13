package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.8 MB,击败了60.32% 的Go用户
func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}
	return true
}
