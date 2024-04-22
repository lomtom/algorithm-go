package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.8 MB,击败了36.79% 的Go用户
func combinationSum(candidates []int, target int) (result [][]int) {
	var dfs func(index, sub int, res []int)
	dfs = func(index, sub int, res []int) {
		if sub == 0 {
			result = append(result, append([]int{}, res...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if sub-candidates[i] >= 0 {
				res = append(res, candidates[i])
				dfs(i, sub-candidates[i], res)
				res = res[:len(res)-1]
			}
		}
	}
	dfs(0, target, []int{})
	return result
}
