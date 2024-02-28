package leetcode

// n = 7, cost = [1,5,2,2,3,3,1]
// 左：2 * i 右：2 * i + 1
func minIncrements(n int, cost []int) int {
	ans := 0
	for i := n - 2; i > 0; i -= 2 {
		ans += abs(cost[i] - cost[i+1])
		// 叶节点 i 和 i+1 的双亲节点下标为 i/2（整数除法）
		cost[i/2] = cost[i/2] + max(cost[i], cost[i+1])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
