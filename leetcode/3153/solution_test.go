package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(sumDigitDifferences([]int{554, 528}))
	t.Log(sumDigitDifferences([]int{824, 843, 837, 620, 836, 234, 276, 859}))
	t.Log(sumDigitDifferences([]int{50, 28, 48}))
	t.Log(sumDigitDifferences([]int{13, 23, 12}))
	t.Log(sumDigitDifferences([]int{10, 10, 10, 10}))
}
