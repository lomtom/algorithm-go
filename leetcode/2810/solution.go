package leetcode

// 执行耗时:3 ms,击败了87.50% 的Go用户
// 内存消耗:3.3 MB,击败了35.00% 的Go用户
func finalString(s string) string {
	var result []byte
	var flag = false
	for index := range s {
		if s[index] == 'i' {
			flag = !flag
		} else {
			if flag {
				result = append(result, s[index])
			} else {
				result = append([]byte{s[index]}, result...)
			}
		}
	}
	if !flag {
		for index := 0; index < len(result)/2; index++ {
			result[index], result[len(result)-index-1] = result[len(result)-index-1], result[index]
		}
	}
	return string(result)
}
