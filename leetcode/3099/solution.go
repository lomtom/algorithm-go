package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了63.60% 的Go用户
func sumOfTheDigitsOfHarshadNumber(x int) int {
	var temp = x
	var sum = 0
	for temp > 0 {
		sum += temp % 10
		temp /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
