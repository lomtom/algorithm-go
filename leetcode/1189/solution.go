package leetcode

var m = map[byte]int{
	'b': 0,
	'a': 1,
	'l': 2,
	'o': 3,
	'n': 4,
}

// 执行耗时:2 ms,击败了33.33% 的Go用户
// 内存消耗:2.3 MB,击败了76.92% 的Go用户
func maxNumberOfBalloons(text string) int {
	count := make([]int, 5)
	for i := 0; i < len(text); i++ {
		if index, ok := m[text[i]]; ok {
			count[index]++
		}
	}
	count[2] /= 2
	count[3] /= 2
	var max = count[0]
	for index := range count {
		if count[index] < max {
			max = count[index]
		}
	}
	return max
}
