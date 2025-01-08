package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.9 MB,击败了53.85% 的Go用户
func largestGoodInteger(num string) string {
	var ans byte = ' '
	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] && num[i] > ans {
			ans = num[i]
		}
	}
	if ans == ' ' {
		return ""
	}
	return string(ans) + string(ans) + string(ans)
}
