package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了100.00% 的Go用户
func timeRequiredToBuy(tickets []int, k int) int {
	var ans int = 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			ans += min(tickets[i], tickets[k])
		} else {
			ans += min(tickets[i], tickets[k]-1)
		}
	}
	return ans
}
