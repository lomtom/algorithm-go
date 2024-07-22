package leetcode

// 执行耗时:34 ms,击败了29.17% 的Go用户
// 内存消耗:5 MB,击败了62.50% 的Go用户
func maximumDetonation(bombs [][]int) int {
	var maxCount int
	var n = len(bombs)
	var connects = make([][]bool, n)
	for i := 0; i < n; i++ {
		connects[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if i != j {
				connects[i][j] = isRange(bombs[i], bombs[j])
			}
		}
	}
	for i := 0; i < n; i++ {
		var visited = make([]bool, n)
		maxCount = max(maxCount, dfs(bombs, visited, connects, i, n))
	}
	return maxCount
}

func dfs(bombs [][]int, visited []bool, connects [][]bool, index, n int) int {
	visited[index] = true
	var count int
	for i := 0; i < n; i++ {
		if !visited[i] && connects[index][i] {
			count += dfs(bombs, visited, connects, i, n)
		}
	}
	return count + 1
}

func isRange(num1, num2 []int) bool {
	return squ(num1[0]-num2[0])+squ(num1[1]-num2[1]) <= squ(num1[2])
}

func squ(num int) int {
	return num * num
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
