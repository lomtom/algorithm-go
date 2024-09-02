package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(maxConsecutiveAnswers("FFFTTFTTFT", 3))
	t.Log(maxConsecutiveAnswers("TTFF", 2))
	t.Log(maxConsecutiveAnswers("TFFT", 1))
	t.Log(maxConsecutiveAnswers("TTFTTFTT", 1))
}
