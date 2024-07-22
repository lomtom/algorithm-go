package leetcode

// 执行耗时:17 ms,击败了13.64% 的Go用户
// 内存消耗:6.4 MB,击败了31.82% 的Go用户
func isBipartite(graph [][]int) bool {
	var parent = make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(i, j int)
	union = func(i, j int) {
		parent[find(i)] = find(j)
	}
	for i := range graph {
		for j := 0; j < len(graph[i]); j++ {
			if find(i) == find(graph[i][j]) {
				return false
			}
			union(graph[i][0], graph[i][j])
		}
	}
	return true
}
