package leetcode

// 执行用时：118 ms, 在所有 Go 提交中击败了64.71%的用户
// 内存消耗：21.15 MB, 在所有 Go 提交中击败了70.%的用户
func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for index := range edges {
		g[edges[index][0]] = append(g[edges[index][0]], edges[index][1])
	}
	var result = make([][]int, n)
	vis := make([]int, n)
	start := 0
	var appendResultFunc func(key int)
	appendResultFunc = func(key int) {
		vis[key] = start + 1 // 避免重复访问
		for _, v := range g[key] {
			if vis[v] != start+1 {
				result[v] = append(result[v], start)
				appendResultFunc(v)
			}
		}
	}
	for ; start < n; start++ {
		appendResultFunc(start) // 从 start 开始 DFS
	}
	return result
}
