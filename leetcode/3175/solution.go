package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:10.9 MB,击败了15.79% 的Go用户
func findWinningPlayer(skills []int, k int) int {
	play1 := skills[0]
	skills = skills[1:]
	var count int
	var num1, num2 = 0, 1
	for count < k {
		play2 := skills[0]
		if play1 > play2 {
			count++
			skills = append(skills[1:], play2)
		} else {
			count = 1
			skills = append(skills[1:], play1)
			play1 = play2
			num1 = num2
		}
		num2++
		if count > len(skills) {
			return num1
		}
	}
	return num1
}
