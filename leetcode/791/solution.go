package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.9 MB,击败了37.50% 的Go用户
func customSortString(order string, s string) string {
	sMap := make(map[uint8]int)
	for index := range s {
		sMap[s[index]] = sMap[s[index]] + 1
	}
	var result []uint8
	for index := range order {
		if v, ok := sMap[order[index]]; ok && v > 0 {
			for i := 0; i < v; i++ {
				result = append(result, order[index])
			}
			sMap[order[index]] = 0
		}
	}
	for k, v := range sMap {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}
	return string(result)
}
