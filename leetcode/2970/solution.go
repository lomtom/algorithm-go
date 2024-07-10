package leetcode

// 执行耗时:2 ms,击败了85.00% 的Go用户
// 内存消耗:2.6 MB,击败了100.00% 的Go用户
func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	res := 0
	l := 1
	for l < n && nums[l-1] < nums[l] {
		l++
	}
	res += l
	if l < n {
		res += 1
	}
	for r := n - 2; r >= 0; r-- {
		for l > 0 && nums[l-1] >= nums[r+1] {
			l--
		}
		res += l
		if l <= r {
			res += 1
		}
		if nums[r] >= nums[r+1] {
			break
		}
	}
	return res
}
