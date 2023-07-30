package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了100.00% 的Go用户
func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	exists := make([]bool, 26)
	for index := range sentence {
		exists[sentence[index]-'a'] = true
	}
	for index := range exists {
		if !exists[index] {
			return false
		}
	}
	return true
}
