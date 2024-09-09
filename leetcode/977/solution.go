package leetcode

// 执行耗时:19 ms,击败了74.61% 的Go用户
// 内存消耗:6.7 MB,击败了39.79% 的Go用户
func sortedSquares(nums []int) []int {
	var index int = 0
	for index < len(nums) && nums[index] < 0 {
		index++
	}
	var left int = index - 1
	var right int = index
	var result []int = make([]int, len(nums))
	index = 0
	for left >= 0 && right < len(nums) {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[index] = nums[right] * nums[right]
			right++
		} else {
			result[index] = nums[left] * nums[left]
			left--
		}
		index++
	}
	for left >= 0 {
		result[index] = nums[left] * nums[left]
		left--
		index++
	}
	for right < len(nums) {
		result[index] = nums[right] * nums[right]
		right++
		index++
	}
	return result
}
