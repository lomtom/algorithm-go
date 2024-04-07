package leetcode

// 执行耗时:6 ms,击败了58.64% 的Go用户
// 内存消耗:4.3 MB,击败了23.06% 的Go用户
func removeDuplicates(nums []int) int {
	var left, right = 0, 0
	for right < len(nums) {
		nums[left] = nums[right]
		for right < len(nums) && nums[left] == nums[right] {
			right++
		}
		left++
	}
	return left
}
