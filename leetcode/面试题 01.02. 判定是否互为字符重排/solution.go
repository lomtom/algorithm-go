package leetcode

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了30.69%的用户
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	byte1 := []byte(s1)
	byte2 := []byte(s2)
	m := make(map[byte]int)
	for index := range byte1 {
		m[byte1[index]]++
	}
	for index := range byte2 {
		if v, ok := m[byte2[index]]; !ok || v <= 0 {
			return false
		} else {
			m[byte2[index]]--
		}
	}
	return true
}
