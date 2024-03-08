package leetcode

import "testing"

func Test(t *testing.T) {
	t.Log(minimumPossibleSum(2, 3))                   // 4: 1 3
	t.Log(minimumPossibleSum(3, 3))                   // 8: 1 3 4
	t.Log(minimumPossibleSum(3, 4))                   // 7: 1 2 4
	t.Log(minimumPossibleSum(1, 1))                   // 1: 1
	t.Log(minimumPossibleSum(39636, 49035))           // 156198582
	t.Log(minimumPossibleSum(1000000000, 1000000000)) // 750000042
	t.Log(minimumPossibleSum(8, 8))                   // 1 2 3 4     8 9 10 11
}
