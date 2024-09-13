package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.4 MB,击败了99.83% 的Go用户
func findMin(nums []int) int {
	var l, r = 0, len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
