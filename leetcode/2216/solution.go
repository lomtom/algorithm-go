package leetcode

// 执行耗时:117 ms,击败了82.61% 的Go用户
// 内存消耗:7.7 MB,击败了100.00% 的Go用户
func minDeletion(nums []int) int {
	var result int
	for i := 0; i < len(nums)-1; i++ {
		for (i-result)%2 == 0 && i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
			result++
		}
	}
	if (len(nums)-result)%2 == 1 {
		result++
	}
	return result
}
