package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.1 MB,击败了6.04% 的Go用户
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var res []float64
	resMap := make(map[string]map[string]float64)
	for i, equation := range equations {
		if _, ok := resMap[equation[0]]; !ok {
			resMap[equation[0]] = make(map[string]float64)
		}
		if _, ok := resMap[equation[1]]; !ok {
			resMap[equation[1]] = make(map[string]float64)
		}
		resMap[equation[0]][equation[1]] = values[i]
		resMap[equation[1]][equation[0]] = 1 / values[i]
	}
	var dfs func(base, to string)
	dfs = func(base, to string) {
		if _, ok := resMap[to]; !ok {
			return
		}
		for key, value := range resMap[to] {
			if _, ok := resMap[base][key]; ok {
				continue
			}
			resMap[base][key] = resMap[base][to] * value
			dfs(base, key)
		}
	}
	for base, pairs := range resMap {
		for key, _ := range pairs {
			dfs(base, key)
		}
	}
	for _, query := range queries {
		if _, ok := resMap[query[0]]; !ok {
			res = append(res, -1.0)
			continue
		}
		if _, ok2 := resMap[query[0]][query[1]]; !ok2 {
			res = append(res, -1.0)
			continue
		}
		res = append(res, resMap[query[0]][query[1]])
	}
	return res
}
