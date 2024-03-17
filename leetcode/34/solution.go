package leetcode

// 执行耗时:6 ms,击败了43.33% 的Go用户
// 内存消耗:4.5 MB,击败了6.44% 的Go用户
func searchRange(nums []int, target int) []int {
	var left, right = 0, len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right = left
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left, right - 1}
}
