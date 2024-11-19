package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:5.3 MB,击败了5.12% 的Go用户
func destCity(paths [][]string) string {
	var m = make(map[string]bool)
	for index := range paths {
		m[paths[index][1]] = true
	}
	for index := range paths {
		delete(m, paths[index][0])
	}
	var res string
	for k, _ := range m {
		res = k
		break
	}
	return res
}
