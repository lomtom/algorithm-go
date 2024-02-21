package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.2 MB,击败了68.75% 的Go用户
func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return
}
