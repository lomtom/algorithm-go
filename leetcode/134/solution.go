package leetcode

// 执行耗时:109 ms,击败了83.13% 的Go用户
// 内存消耗:7.4 MB,击败了99.81% 的Go用户
func canCompleteCircuit(gas []int, cost []int) int {
	result := 0
	max := 0
	index := 0
	for i := len(gas) - 1; i >= 0; i-- {
		result += gas[i] - cost[i]
		if result > max {
			max = result
			index = i
		}
	}
	if result < 0 {
		return -1
	}
	return index
}
