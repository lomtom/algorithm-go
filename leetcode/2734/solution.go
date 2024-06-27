package leetcode

// 执行耗时:31 ms,击败了60.00% 的Go用户
// 内存消耗:8.2 MB,击败了40.00% 的Go用户
func smallestString(s string) string {
	var bytes = []byte(s)
	var flag = true
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 'a' {
			if flag {
				continue
			}
			break
		}
		flag = false
		bytes[i]--
	}
	if flag {
		bytes[len(bytes)-1] = 'z'
	}
	return string(bytes)
}
