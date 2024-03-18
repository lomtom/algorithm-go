package leetcode

// 执行耗时:2 ms,击败了43.29% 的Go用户
// 内存消耗:2.5 MB,击败了5.12% 的Go用户
func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	isLeft := func(i int) bool {
		end := nums[right]
		// 左边为纯单调递增数组
		if nums[i] > end {
			return target > end && nums[i] >= target
		}
		// 右边为纯单调递增数组
		return target > end || nums[i] >= target
	}
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if isLeft(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
