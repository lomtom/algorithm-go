package leetcode

// 执行耗时:49 ms,击败了45.10% 的Go用户
// 内存消耗:9.4 MB,击败了66.67% 的Go用户
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(int, int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	for _, conn := range connections {
		union(conn[0], conn[1])
	}
	var count int
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}
	return count - 1
}
