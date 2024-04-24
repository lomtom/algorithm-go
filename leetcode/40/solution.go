package leetcode

import "sort"

// 执行耗时:2 ms,击败了59.02% 的Go用户
// 内存消耗:2.5 MB,击败了31.62% 的Go用户
func combinationSum2(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	var dfs func(index, sub int, res []int)
	dfs = func(index, sub int, res []int) {
		if sub == 0 {
			result = append(result, append([]int{}, res...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if sub-candidates[i] >= 0 {
				res = append(res, candidates[i])
				dfs(i+1, sub-candidates[i], res)
				res = res[:len(res)-1]
			}
		}
	}
	dfs(0, target, []int{})
	return result
}
