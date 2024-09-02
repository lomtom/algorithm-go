package leetcode

// 执行耗时:4 ms,击败了98.84% 的Go用户
// 内存消耗:5.2 MB,击败了73.26% 的Go用户
func maxConsecutiveAnswers(answerKey string, k int) int {
	var tCount, fCount int
	var left, right = 0, 0
	var ans = 0
	for right < len(answerKey) {
		// 扩张
		if answerKey[right] == 'T' {
			tCount++
		} else {
			fCount++
		}
		right++
		// 收缩
		for tCount > k && fCount > k {
			if answerKey[left] == 'T' {
				tCount--
			} else {
				fCount--
			}
			left++
		}
		ans = max(ans, tCount+fCount)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
