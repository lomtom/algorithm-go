package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了100.00% 的Go用户
func maxValue(n, index, maxSum int) int {
	var left, right = 1, maxSum
	for left < right {
		mid := left + (right-left)>>1
		if mid+cal(mid, index)+cal(mid, n-index-1) < maxSum {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func cal(max, length int) int {
	if length == 0 {
		return 0
	}
	if length <= max {
		return (2*max + 1 - length) * length / 2
	}
	return (max+1)*max/2 + length - max
}
