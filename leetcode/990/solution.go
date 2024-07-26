package leetcode

import "sort"

// 执行耗时:2 ms,击败了50.00% 的Go用户
// 内存消耗:2.6 MB,击败了56.58% 的Go用户
func equationsPossible(equations []string) bool {
	sort.Slice(equations, func(i, j int) bool {
		if equations[i][1] == '!' {
			return false
		}
		return true
	})
	var parent = make([]int, 26)
	for i := 0; i < 26; i++ {
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
	for index := range equations {
		left, symbol, right := equations[index][0], equations[index][1], equations[index][3]
		if symbol == '=' {
			union(int(left-'a'), int(right-'a'))
		} else if find(int(left-'a')) == find(int(right-'a')) {
			return false
		}
	}
	return true
}
