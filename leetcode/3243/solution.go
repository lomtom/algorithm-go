package leetcode

// 执行耗时:463 ms,击败了7.14% 的Go用户
// 内存消耗:8.8 MB,击败了20.83% 的Go用户
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	var dist = make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := i; j < n; j++ {
			dist[i][j] = j - i
		}
	}
	var res = make([]int, len(queries))
	for index := range queries {
		dist[queries[index][0]][queries[index][1]] = 1
		for i := queries[index][0]; i >= 0; i-- {
			for j := queries[index][1]; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][queries[index][0]]+dist[queries[index][0]][j])
				dist[i][j] = min(dist[i][j], dist[i][queries[index][1]]+dist[queries[index][1]][j])
			}
		}
		res[index] = dist[0][n-1]
	}
	return res
}
