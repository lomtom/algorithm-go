package leetcode

// 执行耗时:266 ms,击败了5.00% 的Go用户
// 内存消耗:8.1 MB,击败了60.00% 的Go用户
func minimumLevels(possible []int) int {
	if possible[0] == 0 {
		possible[0] = -1
	}
	for index := 1; index < len(possible); index++ {
		if possible[index] == 0 {
			possible[index] = possible[index-1] - 1
		} else {
			possible[index] = possible[index-1] + 1
		}
	}
	for index := 0; index < len(possible)-1; index++ {
		if possible[len(possible)-1]-possible[index] < possible[index] {
			return index + 1
		}
	}
	return -1
}
