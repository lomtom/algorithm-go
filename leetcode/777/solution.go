package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.7 MB,击败了85.71% 的Go用户
func canTransform(start string, end string) bool {
	var n, startIndex, endIndex = len(start), 0, 0
	for startIndex < n || endIndex < n {
		for startIndex < n && start[startIndex] == 'X' {
			startIndex++
		}
		for endIndex < n && end[endIndex] == 'X' {
			endIndex++
		}
		if startIndex == n || endIndex == n {
			return startIndex == endIndex
		}
		if start[startIndex] != end[endIndex] {
			return false
		}
		if start[startIndex] == 'L' && startIndex < endIndex {
			return false
		}
		if start[startIndex] == 'R' && startIndex > endIndex {
			return false
		}
		startIndex++
		endIndex++
	}
	return startIndex == endIndex
}
