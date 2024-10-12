package leetcode

// 执行耗时:6 ms,击败了17.39% 的Go用户
// 内存消耗:3 MB,击败了43.48% 的Go用户
func duplicateNumbersXOR(nums []int) int {
	var xor int
	var m = make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 2 {
			xor ^= k
		}
	}
	return xor
}
